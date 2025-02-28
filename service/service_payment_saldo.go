package service

import (
	"camera-rent/entity"
	"camera-rent/input"
	"camera-rent/repository"
	"fmt"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var (
	listOfBank = map[string]midtrans.Bank{
		"bca":  midtrans.BankBca,
		"bri":  midtrans.BankBri,
		"cimb": midtrans.BankCimb,
		"bni":  midtrans.BankBni,
	}
)

type midtransGateway struct {
	client *coreapi.Client
}

type Config struct {
	ServerKey string
}

func NewMidtransGateway(cfg *Config) (*midtransGateway, error) {
	client := coreapi.Client{}
	client.New(cfg.ServerKey, midtrans.Sandbox)
	return &midtransGateway{
		client: &client,
	}, nil
}

type ServicePaymentSaldo interface {
	DoPaymentSaldo(req input.SubmitPaymentRequest, topUpId string, userID int) (*entity.DoPayment, error)
	// GetPaymentSaldos() ([]*entity.PaymentSaldo, error)
	// GetPaymentSaldo(ID int) (*entity.PaymentSaldo, error)
	// DeletePaymentSaldo(ID int) error
	HandleNotificationPaymentDonation(req *entity.MidtransNotificationRequest) error
	// FindStatus(orderID string) (*entity.PaymentSaldo, error)
}

type servicePaymentSaldo struct {
	repositoryPaymentSaldo repository.RepositoryPaymentSaldo
	repositoryUser         repository.RepositoryUser
	repositoryTopUp        repository.RepositoryTopUp
	midtransGateway        *midtransGateway
}

func NewServicePaymentSaldo(repositoryPaymentSaldo repository.RepositoryPaymentSaldo, repositoryUser repository.RepositoryUser, repositoryTopUp repository.RepositoryTopUp, midtransGateway *midtransGateway) *servicePaymentSaldo {
	return &servicePaymentSaldo{repositoryPaymentSaldo, repositoryUser, repositoryTopUp, midtransGateway}
}

func (s *servicePaymentSaldo) DoPaymentSaldo(req input.SubmitPaymentRequest, topUpId string, userID int) (*entity.DoPayment, error) {
	topUpIDInt, err := strconv.Atoi(topUpId)
	if err != nil {
		return nil, fmt.Errorf("invalid topUpId: %w", err)
	}

	// Panggil repository dengan topUpID dalam bentuk int
	getPaymentSaldoID, err := s.repositoryTopUp.FindById(topUpIDInt)
	if err != nil {
		return nil, fmt.Errorf("error fetching top-up: %w", err)
	}

	if getPaymentSaldoID.UserID != userID {
		return nil, fmt.Errorf("unauthorized user")
	}

	chosenBank, ok := listOfBank[req.BankTransfer]
	if !ok {
		return nil, fmt.Errorf("unsupported bank")
	}

	resp, err := s.midtransGateway.client.ChargeTransaction(&coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(getPaymentSaldoID.ID),
			GrossAmt: int64(getPaymentSaldoID.Amount),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: chosenBank,
		},
	})

	pay := &entity.PaymentSaldo{
		StatusPayment: "pending",
		TopUpID:       getPaymentSaldoID.ID,
		UserID:        getPaymentSaldoID.UserID,
		TransactionID: resp.TransactionID,
	}

	_, err = s.repositoryPaymentSaldo.Save(pay)
	if err != nil {
		return nil, fmt.Errorf("error saving payment donation: %w", err)
	}

	return mapChargeToResponseDonation(resp)
}

func mapChargeToResponseDonation(resp *coreapi.ChargeResponse) (*entity.DoPayment, error) {
	if resp == nil {
		return nil, fmt.Errorf("nil response received")
	}

	var vaNumbers []entity.VaNumber
	for _, va := range resp.VaNumbers {
		vaNumbers = append(vaNumbers, entity.VaNumber{
			Bank:     va.Bank,
			VaNumber: va.VANumber,
		})
	}

	return &entity.DoPayment{
		TransactionID:   resp.TransactionID,
		OrderID:         resp.OrderID,
		GrossAmount:     resp.GrossAmount,
		VaNumbers:       vaNumbers,
		TransactionTime: resp.TransactionTime,
		MerchantID:      "G260552465",
	}, nil
}

func (s *servicePaymentSaldo) HandleNotificationPaymentDonation(req *entity.MidtransNotificationRequest) error {
	// Cari transaksi pembayaran berdasarkan TransactionID
	findPayment, err := s.repositoryPaymentSaldo.FindByTransactionID(req.TransactionID)
	if err != nil {
		return err
	}

	// Cari user berdasarkan UserID dari transaksi
	user, err := s.repositoryUser.FindById(findPayment.UserID)
	if err != nil {
		return fmt.Errorf("failed to fetch user: %w", err)
	}

	switch req.TransactionStatus {
	case "settlement", "capture":
		findPayment.StatusPayment = "settled"

		user.Saldo += findPayment.TopUps.Amount

		if _, err := s.repositoryUser.Update(user); err != nil {
			return fmt.Errorf("failed to update user saldo: %w", err)
		}

	default:
		findPayment.StatusPayment = req.TransactionStatus
	}

	_, err = s.repositoryPaymentSaldo.Update(findPayment)
	if err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}

	return nil
}
