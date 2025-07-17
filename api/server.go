package api

import (
	"context"
	"database/sql"
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

	router.POST("/api_key", authMiddleware(server.tokenMaker), roleMiddleware("client"), server.CreateKey)
	router.POST("/client", authMiddleware(server.tokenMaker), roleMiddleware("client", "admin"), server.ClientRegister)
	router.POST("/api/shipment", AuthAPIKeyMiddleware(server.store), server.CreateShipment)
	router.POST("/shipment/status", server.UpdateShipmentStatus)

	server.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) generateUniqueShipmentCode(ctx context.Context, length int) (string, error) {
	const maxAttempts = 100
	for i := 0; i < maxAttempts; i++ {
		code, err := util.GenerateRandomCode(length)
		if err != nil {
			return "", err
		}

		exists, err := server.store.CheckShipmentCodeExists(ctx, sql.NullString{String: code, Valid: true})
		if err != nil {
			return "", err
		}

		if !exists {
			return code, nil
		}
	}
	return "", fmt.Errorf("failed to generate unique shipment code after %d attempts", maxAttempts)
}
