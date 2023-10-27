package middleware

import (
	"frame/utils"
	"github.com/gin-gonic/gin"
)

// IsAdminMiddlewareFunc must after auth middleware
var IsAdminMiddlewareFunc = func(context *gin.Context) {
	isAdmin := utils.GetIsAdminFromCtx(context)
	if !isAdmin {
		context.JSON(utils.ErrNoAuth, "no permission")
		context.Abort()
		return
	}
}
