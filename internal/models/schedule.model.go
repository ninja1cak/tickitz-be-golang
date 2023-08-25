package models

import "time"

type Schedule struct {
	Id_schedule    int        `db:"id_schedule" form:"id_schedule" json:"id_schedule,omitempty"`
	Date_start     *time.Time `db:"date_start" form:"date_start" json:"date_start,omitempty"`
	Date_end       *time.Time `db:"date_end" form:"date_end" json:"date_end,omitempty"`
	Cinema_address string     `db:"cinema_address" form:"cinema_address" json:"cinema_address,omitempty"`
	Price_seat     string     `db:"price_seat" form:"price_seat" json:"price_seat,omitempty"`
	City
	Time
	Cinema
}

type City struct {
	Id_city int      `db:"id_city" form:"id_city" json:"id_city,omitempty"`
	City    string   `db:"city" json:"city,omitempty"`
	CityArr []string `form:"city" json:"cityArr,omitempty"`
}

type Time struct {
	Id_time int      `db:"id_time" form:"id_time" json:"id_time,omitempty"`
	Time    string   `db:"time" json:"time,omitempty"`
	TimeArr []string `form:"time" json:"timeArr,omitempty"`
}

type Cinema struct {
	Id_cinema       int      `db:"id_cinema" form:"id_cinema" json:"id_cinema,omitempty"`
	Cinema_name     string   `db:"cinema_name" json:"cinema_name,omitempty"`
	CinemaArr_name  []string `form:"cinema_name" json:"cinemaArr_name,omitempty"`
	Cinema_logo_url string   `db:"cinema_logo_url" form:"cinema_logo_url" json:"cinema_logo_url,omitempty"`
}
