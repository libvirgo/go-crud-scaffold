package user

import (
	"fmt"
	"frame/ent"
	"frame/ent/user"
	"frame/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Register godoc
// @Summary Register
// @Description Register
// @Tags User
// @Accept  json
// @Produce  json
// @Param req body RegisterReq true "wallet address and sign data"
// @Success 200 {object} utils.ResponseVO{data=utils.StatusVO} "code 1001: internal server error, code 1003: invalid params code 1005: invalid sign"
// @Router /user/register [post]
func (h H) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(utils.ErrInvalidParams,
			gin.H{
				"error": "invalid params",
			})
		return
	}
	_, err := h.Ent.User.Query().Where(user.WalletAddressEQ(req.WalletAddress)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		h.Error("query user error", zap.String("wallet_address", req.WalletAddress), zap.String("error", err.Error()))
		ctx.JSON(utils.ErrInternalServer, "query user error")
		return
	}
	if !ent.IsNotFound(err) {
		h.Error("user already exist", zap.String("wallet_address", req.WalletAddress))
		ctx.JSON(utils.ErrInvalidParams, "user already exist")
		return
	}
	nonceKey := h.Redis.Get(ctx, fmt.Sprintf(utils.WalletNonceKey, req.WalletAddress)).Val()
	if nonceKey == "" {
		ctx.JSON(utils.ErrInvalidSign, "nonce is empty, please retry")
		return
	}
	defer func() {
		h.Redis.Del(ctx, fmt.Sprintf(utils.WalletNonceKey, req.WalletAddress))
	}()
	if !utils.VerifySig(req.WalletAddress, req.SignData, []byte(nonceKey)) {
		h.Error("sign data is invalid", zap.String("nonce", nonceKey), zap.String("sign_data", req.SignData))
		ctx.JSON(utils.ErrInvalidSign, "sign data is invalid")
		return
	}
	err = ent.WithTx(ctx, h.Ent, func(tx *ent.Tx) error {
		u, err := tx.User.Create().SetWalletAddress(req.WalletAddress).Save(ctx)
		if err != nil {
			h.Error("create user error", zap.String("wallet_address", req.WalletAddress), zap.Error(err))
			ctx.JSON(utils.ErrInternalServer, "create user error")
			return err
		}
		go func() {
			_, err := h.Ent.UserActivity.Create().SetUser(u).SetType(utils.UserActivityRegister).Save(ctx)
			if err != nil {
				h.Error("create user activity error", zap.String("wallet_address", req.WalletAddress), zap.Error(err))
			}
		}()
		return nil
	})
	if err != nil {
		h.Error("create user error", zap.String("wallet_address", req.WalletAddress), zap.Error(err))
		ctx.JSON(utils.ErrInternalServer, "create user error")
		return
	}
	ctx.JSON(utils.Success, utils.StatusVO{
		Status: "success",
	})
}

type RegisterReq struct {
	WalletAddress string `json:"wallet_address" binding:"required,wallet"`
	SignData      string `json:"sign_data" binding:"required"`
}
