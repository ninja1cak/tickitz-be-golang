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

type HandlerBooking struct {
	*repositories.RepoBooking
}

func NewBooking(r *repositories.RepoBooking) *HandlerBooking {
	return &HandlerBooking{r}
}

func (h *HandlerBooking) PostDataBooking(ctx *gin.Context) {

	var booking models.Booking
	booking.Id_user = ctx.MustGet("user_id").(string)

	if err := ctx.ShouldBind(&booking); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}
	log.Println(booking)
	data, err := h.CreateBooking(&booking)
	if err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return

	} else {
		pkg.NewResponse(200, data).Send(ctx)

	}

}

func (h *HandlerBooking) GetDataBookingByUser(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	id_user := ctx.MustGet("user_id").(string)
	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "3"
	}
	data, err := h.GetBookingByUser(limit, page, id_user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})

	} else {
		pkg.NewResponse(200, data).Send(ctx)
	}

}
