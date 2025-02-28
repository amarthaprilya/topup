package input

import "time"

type BookingInput struct {
	FirstDateRent time.Time `json:"first_date_rent" binding:"required"`
	LastDateRent  time.Time `json:"last_date_rent" binding:"required"`
	ProductID     int       `json:"product_id" binding:"required"`
}
