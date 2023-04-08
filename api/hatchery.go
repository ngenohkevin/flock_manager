package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
	"net/http"
)

type createHatcheryRequest struct {
	ProductionID int64 `json:"production_id" binding:"required"`
	Infertile    int64 `json:"infertile" binding:"required"`
	Early        int64 `json:"early" binding:"required"`
	Middle       int64 `json:"middle" binding:"required"`
	Late         int64 `json:"late" binding:"required"`
	DeadChicks   int64 `json:"dead_chicks" binding:"required"`
	AliveChicks  int64 `json:"alive_chicks" binding:"required"`
}

func (server *Server) createHatchery(ctx *gin.Context) {
	var req createHatcheryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateHatcheryParams{
		ProductionID: req.ProductionID,
		Infertile:    req.Infertile,
		Early:        req.Early,
		Middle:       req.Middle,
		Late:         req.Late,
		DeadChicks:   req.DeadChicks,
		AliveChicks:  req.AliveChicks,
	}
	hatchery, err := server.store.CreateHatchery(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, hatchery)
}
