package handlers

import (
	"net/http"
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/repositories"
	"ninja1cak/coffeshop-be/pkg"

	"github.com/gin-gonic/gin"
)

type HandleSchedule struct {
	*repositories.RepoSchedule
}

func NewSchedule(r *repositories.RepoSchedule) *HandleSchedule {
	return &HandleSchedule{r}
}

func (h *HandleSchedule) GetDataSchedule(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	location := ctx.Query("location")
	time := ctx.Query("time")
	id_movie := ctx.Query("id_movie")
	date := ctx.Query("date")
	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "3"
	}
	data, err := h.GetSchedule(limit, page, location, time, id_movie, date)
	if err != nil {

		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)

	} else {
		pkg.NewResponse(200, data).Send(ctx)
	}

}

func (h *HandleSchedule) GetDataCity(ctx *gin.Context) {

	data, err := h.GetCity()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})

	} else {
		pkg.NewResponse(200, data).Send(ctx)
	}

}

func (h *HandleSchedule) GetDataTime(ctx *gin.Context) {

	data, err := h.GetTime()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})

	} else {
		pkg.NewResponse(200, data).Send(ctx)
	}

}

func (h *HandleSchedule) GetDataCinema(ctx *gin.Context) {

	data, err := h.GetCinema()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})

	} else {
		pkg.NewResponse(200, data).Send(ctx)
	}

}
