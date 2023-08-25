package routers

import (
	"ninja1cak/coffeshop-be/internal/handlers"
	"ninja1cak/coffeshop-be/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/")

	//dependency injection
	repo := repositories.NewUser(d)
	handler := handlers.NewAuth(repo)

	router.POST("/login", handler.Login)
	router.GET("/auth/:token", handler.VerifyAccount)

}
