package user

import (
	"fmt"
	"frame/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

// GetLoginNonce godoc
// @Summary Get login nonce
// @Description Get login nonce
// @Tags User
// @Accept  json
// @Produce  json
// @Param wallet_address query string true "wallet address"
// @Success 200 {object} utils.ResponseVO{data=NonceResp} "code 1003: invalid params"
// @Router /user/login/nonce [get]
func (h H) GetLoginNonce(ctx *gin.Context) {
	var req NonceReq
	if err := ctx.ShouldBindWith(&req, binding.Query); err != nil {
		ctx.JSON(utils.ErrInvalidParams,
			gin.H{
				"error": err.Error(),
			})
		return
	}
	nonce := fmt.Sprintf("welcome login nonce:%d", time.Now().Unix())
	h.Redis.SetEx(ctx, fmt.Sprintf(utils.WalletLoginNonceKey, req.WalletAddress), nonce, time.Hour*24)
	ctx.JSON(utils.Success, NonceResp{
		Nonce: nonce,
	})
}
