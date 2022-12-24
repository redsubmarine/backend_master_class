package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/redsubmarine/simplebank/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	fmt.Println("17", ctx.Params, ctx.Keys)
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("21", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println("23")
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	fmt.Println("31")
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		fmt.Println("35")
		return
	}
	fmt.Println("38")
	ctx.JSON(http.StatusOK, account)
}
