package user

import (
	"fmt"
	"frame/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// Healthy godoc
// @Summary Healthy
// @Description Healthy
// @Tags User
// @Accept  json
// @Produce  json
// @Param wallet_address query string false "wallet address"
// @Success 200 {object} utils.ResponseVO{data=HealthyResp} "code 1003: invalid params"
// @Router /user/healthy [get]
func (h H) Healthy(ctx *gin.Context) {
	fmt.Println("Healthy", time.Now().Format("2006-01-02 15:04:05"))
	ctx.JSON(utils.Success, HealthyResp{
		Version: "v1.0.2",
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	})
}

type HealthyResp struct {
	Version string `json:"version"`
	Time    string `json:"time"`
}
