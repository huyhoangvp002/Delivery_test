package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/huyhoangvp002/Delivery_test/db/sqlc"
	"github.com/huyhoangvp002/Delivery_test/token"
	"github.com/huyhoangvp002/Delivery_test/util"
)

type Server struct {
	store      db.Querier
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Querier) (*Server, error) {

	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setUpRouter()
	fmt.Println("[DEBUG] TimeOut:", server.config.AccessTokenDuration)
	fmt.Println("[DEBUG] Server:", server.config.ServerAddress)

	return server, nil
}

func (server *Server) setUpRouter() {

	router := gin.Default()

	router.POST("/signup", server.CreateAccount)
	router.POST("/signin", server.Login)

	server.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
