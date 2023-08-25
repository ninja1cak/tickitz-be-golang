package pkg

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Server(router *gin.Engine) *http.Server {
	var address string = "0.0.0.0:8081"
	if port := os.Getenv("PORT"); port != "" {
		address = ":" + port
	}
	corss := cors.AllowAll()
	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      corss.Handler(router),
	}
	return srv
}
