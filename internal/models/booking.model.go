package models

type Booking struct {
	Id_booking           int    `db:"id_booking" form:"id_booking" json:"id_booking,omitempty"`
	Seats_booking        string `db:"seats_booking" form:"seats_booking" json:"seats_booking,omitempty"`
	Total_prices_booking int    `db:"total_prices_booking" form:"total_prices_booking" json:"totals_price_booking,omitempty"`
	Watch_date           string `db:"watch_date" form:"watch_date" json:"watch_date,omitempty"`
	Watch_time           string `db:"watch_time" form:"watch_time" json:"watch_time,omitempty"`
	Payment_method       string `db:"payment_method" form:"payment_method" json:"payment_method,omitempty"`
	User
	Schedule
	Cinema
	Movie
}
