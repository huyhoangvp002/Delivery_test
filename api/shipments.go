package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/huyhoangvp002/Delivery_test/db/sqlc"
)

type addressRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type shipmentRequest struct {
	FromAddress addressRequest `json:"from_address" binding:"required"`
	ToAddress   addressRequest `json:"to_address" binding:"required"`
	Fee         int            `json:"fee" binding:"required,min=0"`
}

func (server *Server) CreateShipment(ctx *gin.Context) {
	var req shipmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	clientIDValue, exists := ctx.Get("client_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	clientID := clientIDValue.(int64)

	arg := db.CreateAddressParams{
		Name:    req.FromAddress.Name,
		Phone:   req.FromAddress.Phone,
		Address: req.FromAddress.Address,
		Status:  "from",
	}

	from_address, err := server.store.CreateAddress(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg = db.CreateAddressParams{
		Name:    req.ToAddress.Name,
		Phone:   req.ToAddress.Phone,
		Address: req.ToAddress.Address,
		Status:  "to",
	}

	to_address, err := server.store.CreateAddress(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	shipment_code, err := server.generateUniqueShipmentCode(ctx, 8)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg1 := db.CreateShipmentParams{
		ClientID:      sql.NullInt64{Int64: clientID, Valid: true},
		FromAddressID: sql.NullInt64{Int64: from_address.ID, Valid: true},
		ToAddressID:   sql.NullInt64{Int64: to_address.ID, Valid: true},
		Fee:           int32(req.Fee),
		ShipmentCode:  sql.NullString{String: shipment_code, Valid: true},
		Status:        sql.NullString{String: "pending", Valid: true},
	}

	shipment, err := server.store.CreateShipment(ctx, arg1)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, shipment)

}
