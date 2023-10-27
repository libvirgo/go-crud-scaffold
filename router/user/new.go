package user

import (
	"frame/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetNonce(ctx *gin.Context)
	Register(ctx *gin.Context)
	GetInfo(ctx *gin.Context)
	GetLoginNonce(ctx *gin.Context)
}

type H utils.Handler

func NewUserRouter(api *gin.RouterGroup, handler utils.Handler,
	authMiddleware *jwt.GinJWTMiddleware,
) utils.Router {
	h := H(handler)
	{
		api.POST("/login", authMiddleware.LoginHandler)
		api.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	g := api.Group("/user")
	{
		g.POST("/register", h.Register)
		g.GET("/nonce", h.GetNonce)
		g.GET("/login/nonce", h.GetLoginNonce)
		g.GET("/info", authMiddleware.MiddlewareFunc(), h.GetInfo)
	}
	return utils.Void{}
}
