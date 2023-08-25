package routers

import (
	"ninja1cak/coffeshop-be/internal/handlers"
	"ninja1cak/coffeshop-be/internal/middleware"
	"ninja1cak/coffeshop-be/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func booking(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/booking")

	//dependency injection
	repo := repositories.NewBooking(d)
	handler := handlers.NewBooking(repo)

	router.POST("/", middleware.IsVerify("admin", "user"), handler.PostDataBooking)
	router.GET("/", middleware.IsVerify("admin", "user"), handler.GetDataBookingByUser)

}
