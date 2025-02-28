package input

type InputTopUp struct {
	Amount int `json:"amount" binding:"required"`
}
