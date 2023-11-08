package user_handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	db "github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users"
	"github.com/madalingrecescu/PizzaDelivery/internal/token"
	"github.com/madalingrecescu/PizzaDelivery/internal/util"
	"net/http"
	"strings"
)

type createAccountRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

type accountResponse struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

type userResponse struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
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
		Username:    account.Username,
		Email:       account.Email,
		PhoneNumber: account.PhoneNumber,
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
		Username:    account.Username,
		Email:       account.Email,
		PhoneNumber: account.PhoneNumber,
	}

	authPayload := ctx.MustGet("authorization_payload").(*token.Payload)
	if authPayload.Username != account.Username {
		ctx.JSON(http.StatusUnauthorized, errors.New(http.StatusUnauthorized, "You are not logged into this account"))
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetAccountByUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, errors.New(401, "Unauthorized. Invalid credentials"))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errors.New(401, "Unauthorized. Invalid credentials"))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getUserFromHeader(ctx *gin.Context) {

	// Get the Authorization header
	headerValue := ctx.GetHeader("Authorization")
	// Extract the token part after "Bearer "
	tokenParts := strings.Split(headerValue, "Bearer ")
	if len(tokenParts) != 2 {
		// Invalid token format
		ctx.JSON(http.StatusBadRequest, errors.New(http.StatusBadRequest, "Invalid Bearer token format"))
		return
	}
	accessToken := tokenParts[1]

	// Proceed with token verification and user retrieval
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errors.New(http.StatusUnauthorized, "Invalid or expired access token"))
		return
	}

	// Use the payload to retrieve the user details from the database
	user, err := server.store.GetAccountByUsername(ctx, payload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := userResponse{
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	ctx.JSON(http.StatusOK, rsp)

}
