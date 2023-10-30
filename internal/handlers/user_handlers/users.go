package user_handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
	"net/http"
)

type createAccountRequest struct {
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	HashedPassword string `json:"hashedPassword" binding:"required,min=6"`
	PhoneNumber    string `json:"phoneNumber" binding:"required"`
}

type accountResponse struct {
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	HashedPassword string `json:"-" binding:"required,min=6"`
	PhoneNumber    string `json:"phoneNumber" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		PhoneNumber:    req.PhoneNumber,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := accountResponse{
		Username:       account.Username,
		Email:          account.Email,
		HashedPassword: account.HashedPassword,
		PhoneNumber:    account.PhoneNumber,
	}
	ctx.JSON(http.StatusOK, rsp)
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
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := accountResponse{
		Username:       account.Username,
		Email:          account.Email,
		HashedPassword: account.HashedPassword,
		PhoneNumber:    account.PhoneNumber,
	}
	ctx.JSON(http.StatusOK, rsp)

}
