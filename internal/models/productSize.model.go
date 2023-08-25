package models

import "time"

type Product_size struct {
	Psize_id          string     `db:"psize_id" form:"psize_id" json:"psize_id"`
	Product_id        string     `db:"product_id" form:"product_id"`
	Product_size      string     `db:"product_size"`
	ProductSizeSlice  []string   `form:"product_size"`
	Product_price     int        `db:"product_price"`
	ProductPriceSlice []int      `form:"product_price"`
	updated_at        *time.Time `db:"updated_at" form:"updated_at"`
}
