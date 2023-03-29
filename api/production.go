package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
	"net/http"
)

type createProductionRequest struct {
	BreedID      int64 `json:"breed_id" binding:"required"`
	Eggs         int64 `json:"eggs" binding:"required"`
	Dirty        int64 `json:"dirty" binding:"required"`
	WrongShape   int64 `json:"wrong_shape" binding:"required"`
	WeakShell    int64 `json:"weak_shell" binding:"required"`
	Damaged      int64 `json:"damaged" binding:"required"`
	HatchingEggs int64 `json:"hatching_eggs" binding:"required"`
}

func (server *Server) createProduction(ctx *gin.Context) {
	var req createProductionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateProductionParams{
		BreedID:      req.BreedID,
		Eggs:         req.Eggs,
		Dirty:        req.Dirty,
		WrongShape:   req.WrongShape,
		WeakShell:    req.WeakShell,
		Damaged:      req.Damaged,
		HatchingEggs: req.HatchingEggs,
	}
	prod, err := server.store.CreateProduction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, prod)
}

type getProductionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProduction(ctx *gin.Context) {
	var req getProductionRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	prod, err := server.store.GetProduction(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, prod)
}