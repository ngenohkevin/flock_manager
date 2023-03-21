package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/breed", server.createBreed)
	router.GET("/breed/:id", server.getBreed)
	router.GET("/breeds", server.listBreeds)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}