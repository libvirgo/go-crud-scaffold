package utils

import (
	"github.com/gin-gonic/gin"
)

const (
	GinPayloadCtxKey       = "payload"
	GinWalletAddressCtxKey = "wallet_address"
	GinUserIdCtxKey        = "uid"
	GinIsAdminCtxKey       = "is_admin"
)

type GinPayload struct {
	WalletAddress string
	IsAdmin       bool
	UserId        int
}

func GetWalletFromCtx(ctx *gin.Context) string {
	data, ok := ctx.Get(GinPayloadCtxKey)
	if !ok {
		return ""
	}
	p, ok := data.(*GinPayload)
	if !ok {
		return ""
	}
	return p.WalletAddress
}

func GetUserIdFromCtx(ctx *gin.Context) int {
	data, ok := ctx.Get(GinPayloadCtxKey)
	if !ok {
		return 0
	}
	p, ok := data.(*GinPayload)
	if !ok {
		return 0
	}
	return p.UserId
}

func GetIsAdminFromCtx(ctx *gin.Context) bool {
	data, ok := ctx.Get(GinPayloadCtxKey)
	if !ok {
		return false
	}
	p, ok := data.(*GinPayload)
	if !ok {
		return false
	}
	return p.IsAdmin
}
