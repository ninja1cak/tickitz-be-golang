package routers

import (
	"ninja1cak/coffeshop-be/internal/handlers"
	"ninja1cak/coffeshop-be/internal/middleware"
	"ninja1cak/coffeshop-be/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/product")

	//dependency injection
	repo := repositories.NewProduct(d)
	handler := handlers.NewProduct(repo)

	router.POST("/", middleware.IsVerify("admin"), middleware.UploadFile, handler.PostDataProduct)
	router.GET("/", handler.GetDataProduct)

	router.PATCH("/:product_slug", middleware.IsVerify("admin"), middleware.UploadFile, handler.UpdateDataProduct)
	router.DELETE("/:product_slug", middleware.IsVerify("admin"), handler.DeleteDataProduct)

}
