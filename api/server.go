package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
	"github.com/ngenohkevin/flock_manager/db/util"
	"github.com/ngenohkevin/flock_manager/token"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create a token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setUpRouter()

	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()
	//user Route
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

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
	router.GET("/hatchery/:id", server.getHatchery)
	router.GET("/hatchery", server.listHatchery)
	router.PUT("/hatchery/:id", server.updateHatchery)
	router.DELETE("/hatchery/:id", server.deleteHatchery)

	//Premises Routes
	router.POST("/premise", server.createPremise)
	router.GET("/premise/:id", server.getPremise)
	router.GET("/premise", server.listPremises)
	router.PUT("/premise/:id", server.updatePremise)
	router.DELETE("/premise/:id", server.deletePremise)

	server.router = router
}

// Start Run the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errors
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
