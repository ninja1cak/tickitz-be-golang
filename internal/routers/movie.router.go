package routers

import (
	"ninja1cak/coffeshop-be/internal/handlers"
	"ninja1cak/coffeshop-be/internal/middleware"
	"ninja1cak/coffeshop-be/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func movie(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/movie")

	//dependency injection
	repo := repositories.NewMovie(d)
	handler := handlers.NewMovie(repo)

	router.POST("/", middleware.IsVerify("user", "admin"), middleware.UploadFile, handler.PostDataMovie)
	router.GET("/", handler.GetDataMovie)
	router.PATCH("/:id_movie", middleware.IsVerify("user", "admin"), middleware.UploadFile, handler.UpdateDatamovie)
	router.DELETE("/:id_movie", middleware.IsVerify("user", "admin"), handler.DeleteDatamovie)

}
