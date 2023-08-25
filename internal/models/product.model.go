package models

import "time"

type Product struct {
	Product_id          string     `db:"product_id" form:"product_id"`
	Product_name        string     `db:"product_name" form:"product_name"`
	Product_desc        string     `db:"product_desc" form:"product_desc"`
	Product_stock       int        `db:"product_stock" form:"product_stock"`
	Product_type        string     `db:"product_type" form:"product_type" json:"product_type"`
	Product_slug        string     `db:"product_slug" form:"product_slug"`
	Product_image       *string    `db:"product_image" form:"product_image"`
	Isfavorite          bool       `db:"isfavorite" form:"isFavorite"`
	Delivery_method     string     `db:"delivery_method" form:"delivery_method"`
	Delivery_hour_start string     `db:"delivery_hour_start" form:"delivery_hour_start"`
	Delivery_hour_end   string     `db:"delivery_hour_end" form:"delivery_hour_end"`
	Created_at          *time.Time `db:"created_at" form:"created_at" json:"created_at"`
	Updated_at          *time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
	Size_price          string     `db:"size_price"`
	Product_size
}
