package routers

import (
	"ninja1cak/coffeshop-be/internal/handlers"
	"ninja1cak/coffeshop-be/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func schedule(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/schedule")

	//dependency injection
	repo := repositories.NewSchedule(d)
	handler := handlers.NewSchedule(repo)

	router.GET("/", handler.GetDataSchedule)
	router.GET("/city", handler.GetDataCity)
	router.GET("/time", handler.GetDataTime)
	router.GET("/cinema", handler.GetDataCinema)

}
