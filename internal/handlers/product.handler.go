package handlers

import (
	"log"
	"net/http"
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/models"
	"ninja1cak/coffeshop-be/internal/repositories"
	"ninja1cak/coffeshop-be/pkg"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	*repositories.RepoProduct
}

func NewProduct(r *repositories.RepoProduct) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) PostDataProduct(ctx *gin.Context) {

	var product models.Product
	var productSize models.Product_size
	product.Product_image = ctx.MustGet("image").(*string)
	if err := ctx.ShouldBind(&product); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}

	if err := ctx.ShouldBind(&productSize); err != nil {
		log.Println("tessssssssssssss", productSize)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}

	data, err := h.CreateProduct(&product, &productSize)
	if err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return

	} else {
		pkg.NewResponse(200, data).Send(ctx)

	}

}

func (h *HandlerProduct) GetDataProduct(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	search := ctx.Query("search")
	sort := ctx.Query("sort")

	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "3"
	}
	data, err := h.GetProduct(limit, page, search, sort)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})

	} else {
		pkg.NewResponse(200, data).Send(ctx)
	}

}

func (h *HandlerProduct) UpdateDataProduct(ctx *gin.Context) {
	var product models.Product
	var productSize models.Product_size
	product.Product_image = ctx.MustGet("image").(*string)

	product.Product_slug = ctx.Param("product_slug")
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}

	if err := ctx.ShouldBind(&productSize); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}

	response, err := h.UpdateProduct(&product, &productSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": "Updated",
			"data":    response,
		})
	}

}

func (h *HandlerProduct) DeleteDataProduct(ctx *gin.Context) {
	var product models.Product

	product.Product_slug = ctx.Param("product_slug")

	response, err := h.DeleteProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": "Deleted",
			"data":    response,
		})
	}
}
