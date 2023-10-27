package user

import (
	"frame/ent/user"
	"frame/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetInfo godoc
// @Summary GetInfo
// @Description GetInfo
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.ResponseVO{data=GetInfoResp} "code 1001: internal server error"
// @Router /user/info [get]
// @Security Bearer
func (h H) GetInfo(ctx *gin.Context) {
	wallet := utils.GetWalletFromCtx(ctx)
	info, err := h.Ent.User.Query().Where(user.WalletAddressEQ(wallet)).First(ctx)
	if err != nil {
		h.Error("query user error", zap.String("wallet_address", wallet), zap.String("error", err.Error()))
		ctx.JSON(utils.ErrInternalServer, "query user error")
		return
	}

	ctx.JSON(utils.Success, GetInfoResp{
		WalletAddress: info.WalletAddress,
		IsAdmin:       utils.GetIsAdminFromCtx(ctx),
	})
}

type GetInfoResp struct {
	WalletAddress string `json:"wallet_address"`
	IsAdmin       bool   `json:"is_admin"`
}
