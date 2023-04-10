package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
	"net/http"
	"strconv"
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

type getPremisesRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPremise(ctx *gin.Context) {
	var req getPremisesRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	premise, err := server.store.GetPremises(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, premise)

}

type listPremisesRequest struct {
	ID       int64 `json:"id"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPremises(ctx *gin.Context) {
	var req listPremisesRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListPremisesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	premise, err := server.store.ListPremises(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, premise)
}

type updatePremisesRequest struct {
	Farm  string `json:"farm" binding:"required"`
	House string `json:"house" binding:"required"`
}

func (server *Server) updatePremise(ctx *gin.Context) {
	var req updatePremisesRequest

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePremisesParams{
		ID:    id,
		Farm:  req.Farm,
		House: req.House,
	}

	premise, err := server.store.UpdatePremises(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, premise)

}
