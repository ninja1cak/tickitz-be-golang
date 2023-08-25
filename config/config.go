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
	AllowOrigins:     []string{"*", "http://localhost:3000"},
	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "HEAD"},
	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
}
