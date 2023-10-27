package user

import (
	"fmt"
	"frame/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

// GetNonce godoc
// @Summary Get nonce
// @Description Get nonce
// @Tags User
// @Accept  json
// @Produce  json
// @Param wallet_address query string true "wallet address"
// @Success 200 {object} utils.ResponseVO{data=NonceResp} "code 1003: invalid params"
// @Router /user/nonce [get]
func (h H) GetNonce(ctx *gin.Context) {
	var req NonceReq
	if err := ctx.ShouldBindWith(&req, binding.Query); err != nil {
		ctx.JSON(utils.ErrInvalidParams,
			gin.H{
				"error": err.Error(),
			})
		return
	}
	nonce := fmt.Sprintf("welcome nonce:%d", time.Now().Unix())
	h.Redis.SetEx(ctx, fmt.Sprintf(utils.WalletNonceKey, req.WalletAddress), nonce, time.Hour*24)
	ctx.JSON(utils.Success, NonceResp{
		Nonce: nonce,
	})
}

type NonceReq struct {
	WalletAddress string `form:"wallet_address" binding:"required,wallet"`
}

type NonceResp struct {
	Nonce string `json:"nonce"`
}
