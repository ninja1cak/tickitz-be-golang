package config

import (
	"github.com/gin-contrib/cors"
)

type Meta struct {
	Next  interface{} `json:"next"`
	Prev  interface{} `json:"prev"`
	Total int         `json:"total"`
}

type Result struct {
	Data    interface{}
	Meta    interface{}
	Message interface{}
}

var CorsConfig = cors.Config{
	AllowOrigins:     []string{"https://foo.com", "*"},
	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
	AllowHeaders:     []string{"Origin", "Authorization"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
}
