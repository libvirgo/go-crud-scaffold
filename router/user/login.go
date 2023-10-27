package user

import "github.com/gin-gonic/gin"

// Login godoc
// @Summary Login
// @Description Login
// @Tags User
// @Accept  json
// @Produce  json
// @Param req body LoginReq true "wallet address"
// @Success 200 {object} utils.ResponseVO{data=LoginResp}
// @Router /login [post]
func Login(ctx *gin.Context) {

}

type LoginReq struct {
	WalletAddress string `json:"wallet_address"`
	SignData      string `json:"sign_data"`
}

type LoginResp struct {
	Code   string `json:"code"`
	Expire string `json:"expire"` // "2023-10-23T19:38:35+08:00"
	Token  string `json:"token"`
}
