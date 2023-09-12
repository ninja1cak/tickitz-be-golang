package handlers

import (
	"log"
	"net/http"
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/models"
	"ninja1cak/coffeshop-be/internal/repositories"
	"ninja1cak/coffeshop-be/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerMovie struct {
	*repositories.RepoMovie
}

func NewMovie(r *repositories.RepoMovie) *HandlerMovie {
	return &HandlerMovie{r}
}

func (h *HandlerMovie) PostDataMovie(ctx *gin.Context) {

	var movie models.Movie
	movie.Url_image_movie = ctx.MustGet("image").(*string)
	if err := ctx.ShouldBind(&movie); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}

	data, err := h.CreateMovie(&movie)
	if err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return

	} else {
		pkg.NewResponse(200, data).Send(ctx)

	}

}

func (h *HandlerMovie) GetDataMovie(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	search := ctx.Query("search")
	sort := ctx.Query("sort")
	searchByIdMovie := ctx.Query("id_movie")
	releaseDateMovie := ctx.Query("release_date")
	log.Println("release date", releaseDateMovie)
	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "3"
	}
	data, err := h.GetMovie(limit, page, search, sort, searchByIdMovie, releaseDateMovie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})

	} else {
		pkg.NewResponse(200, data).Send(ctx)
	}

}

func (h *HandlerMovie) GetDataGenre(ctx *gin.Context) {

	data, err := h.GetGenre()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})

	} else {
		pkg.NewResponse(200, data).Send(ctx)
	}

}

func (h *HandlerMovie) UpdateDatamovie(ctx *gin.Context) {
	var movie models.Movie
	movie.Url_image_movie = ctx.MustGet("image").(*string)

	id_movie, err := strconv.Atoi(ctx.Param("id_movie"))
	movie.Id_movie = id_movie
	if err := ctx.ShouldBind(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}

	response, err := h.UpdateMovie(&movie)
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

func (h *HandlerMovie) DeleteDatamovie(ctx *gin.Context) {
	var movie models.Movie

	id_movie, err := strconv.Atoi(ctx.Param("id_movie"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ctx.Error(err),
		})
		return
	}
	movie.Id_movie = id_movie
	response, err := h.DeleteMovie(&movie)
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
