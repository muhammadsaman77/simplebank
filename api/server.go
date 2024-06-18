package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/muhammadsaman77/simplebank/db/sqlc"
	"github.com/muhammadsaman77/simplebank/token"
	"github.com/muhammadsaman77/simplebank/util"
)



type Server struct {
	config util.Config
	store db.Store
	tokenMaker token.Maker
	router *gin.Engine

	
}

func NewServer(config util.Config, store db.Store)( *Server ,error){
	tokenMaker,err:= token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{config:config, store: store,tokenMaker: tokenMaker}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency", validCurrency)

	}
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.getAccounts)
	router.POST("/transfers/", server.createTransfer)
	router.POST("/users", server.createUser)
	server.router = router
	return server,nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}