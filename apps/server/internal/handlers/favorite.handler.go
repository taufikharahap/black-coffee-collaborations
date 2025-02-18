package handlers

import (
	"strconv"

	"github.com/Roisfaozi/black-coffee-collaborations/config"
	"github.com/Roisfaozi/black-coffee-collaborations/internal/models/favorite"
	"github.com/Roisfaozi/black-coffee-collaborations/internal/repository"
	"github.com/Roisfaozi/black-coffee-collaborations/pkg"

	"github.com/gin-gonic/gin"
)

type HandlerFavorite struct {
	repository.RepoFavoriteIF
}

func NewFavorite(r repository.RepoFavoriteIF) *HandlerFavorite {
	return &HandlerFavorite{r}
}

// Create Favorite
func (h *HandlerFavorite) PostFavorite(ctx *gin.Context) {
	favorite := favorite.FavoriteData{}
	if err := ctx.ShouldBind(&favorite); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	result, err := h.CreateFavorite(&favorite)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(201, result).Send(ctx)
}

// Get Favorite
func (h *HandlerFavorite) GetFavorite(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	offset := (page - 1) * 10

	result, err := h.FetchFavorite(user_id, page, offset)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}

// Delete Favorite
func (h *HandlerFavorite) DeleteFavorite(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	product_id := ctx.Param("product_id")

	result, err := h.RemoveFavorite(user_id, product_id)
	if err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}
