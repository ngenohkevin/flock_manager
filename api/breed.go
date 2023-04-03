package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
	"net/http"
	"strconv"
)

type createBreedRequest struct {
	Breed string `json:"breed" binding:"required"`
}

// createBreed Handler
func (server *Server) createBreed(ctx *gin.Context) {
	var req createBreedRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	breed, err := server.store.CreateBreed(ctx, req.Breed)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, breed)
}

type getBreedRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getBreed handler
func (server *Server) getBreed(ctx *gin.Context) {
	var req getBreedRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	breed, err := server.store.GetBreed(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, breed)
}

type listBreedRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// ListBreeds handler
func (server *Server) listBreeds(ctx *gin.Context) {
	var req listBreedRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListBreedsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	breeds, err := server.store.ListBreeds(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, breeds)
}

type updateBreedRequest struct {
	BreedName string `json:"breed_name" binding:"required"`
}

// update breeds handler
func (server *Server) updateBreed(ctx *gin.Context) {
	var req updateBreedRequest

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateBreedParams{
		BreedID:   id,
		BreedName: req.BreedName,
	}
	breeds, err := server.store.UpdateBreed(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, breeds)
}

type deleteBreedsRequest struct {
	ID      int64 `uri:"id" binding:"required,min=1"`
	BreedID int64 `uri:"id" binding:"required,min=1"`
}

// deleteBreed handler
func (server *Server) deleteBreed(ctx *gin.Context) {
	var req deleteBreedsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// Delete production row first before proceeding to delete breed
	err := server.store.DeleteProduction(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//Show if there's any breed before deleting
	_, err = server.store.GetBreed(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//Delete the breed
	err = server.store.DeleteBreed(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "Deleted",
	})
}
