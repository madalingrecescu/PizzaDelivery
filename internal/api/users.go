package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	db "pizzeria/internal/sqlc_users"
)

type createAccountRequest struct {
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required"`
	HashedPassword string `json:"hashedPassword" binding:"required"`
	PhoneNumber    string `json:"phoneNumber" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
		PhoneNumber:    req.PhoneNumber,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	UserId int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := server.store.GetAccount(ctx, req.UserId)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}
