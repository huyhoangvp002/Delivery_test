package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/huyhoangvp002/Delivery_test/db/sqlc"
	"github.com/huyhoangvp002/Delivery_test/token"
)

func (server *Server) CreateKey(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	account_id_raw, err := server.store.GetAccountIDByUsername(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	account_id := sql.NullInt32{
		Int32: int32(account_id_raw),
		Valid: true,
	}

	client_id_raw, err := server.store.GetClientIDByAccountID(ctx, account_id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	client_id := sql.NullInt64{
		Int64: client_id_raw,
		Valid: true,
	}

	arg := db.CreateApiKeyParams{
		ClientID: client_id,
	}
}
