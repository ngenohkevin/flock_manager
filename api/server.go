package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()
	//Breeds route
	router.POST("/breeds", server.createBreed)
	router.GET("/breeds/:id", server.getBreed)
	router.GET("/breeds", server.listBreeds)
	router.PUT("/breeds/:id", server.updateBreed)
	router.DELETE("/breeds/:id", server.deleteBreed)

	//Production Routes
	router.POST("/production", server.createProduction)
	router.GET("/production/:id", server.getProduction)
	router.GET("/production", server.listProduction)
	router.PUT("/production/:id", server.updateProduction)
	router.DELETE("/production/:id", server.deleteProduction)

	//Hatchery Routes
	router.POST("/hatchery", server.createHatchery)
	//router.GET("/hatchery/:id", server.getProduction)
	//router.GET("/hatchery", server.listProduction)
	//router.PUT("/hatchery/:id", server.updateProduction)
	//router.DELETE("/hatchery/:id", server.deleteProduction)

	server.router = router
	return server
}

// Start Run the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errors
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
