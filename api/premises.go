package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
	"net/http"
)

type createPremisesRequest struct {
	BreedID int64  `json:"breed_id" binding:"required"`
	Farm    string `json:"farm" binding:"required"`
	House   string `json:"house" binding:"required"`
}

func (server *Server) createPremise(ctx *gin.Context) {
	var req createPremisesRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePremisesParams{
		BreedID: req.BreedID,
		Farm:    req.Farm,
		House:   req.House,
	}

	premise, err := server.store.CreatePremises(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, premise)
}
