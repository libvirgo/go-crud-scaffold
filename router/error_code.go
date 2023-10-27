package router

import (
	"frame/utils"
	"github.com/gin-gonic/gin"
)

// GetAllErrorCode godoc
// @Summary Get All Error Code
// @Description Get All Error Code
// @Description 1001: internal server error
// @Description 1002: data not found
// @Description 1003: invalid params
// @Description 1004: no auth
// @Description 1005: invalid blockchain sign data
// @Tags Error Code
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.ResponseVO{data=GetAllErrorCodeResp}
// @Router /errorcode [get]
// @Security Bearer
func GetAllErrorCode(ctx *gin.Context) {
	ctx.JSON(utils.Success, GetAllErrorCodeResp{
		Code: map[int]string{
			utils.ErrInternalServer: "internal server error",
			utils.ErrDataNotFound:   "data not found",
			utils.ErrInvalidParams:  "invalid params",
			utils.ErrNoAuth:         "no auth",
			utils.ErrInvalidSign:    "invalid blockchain sign data",
		},
	})
}

type GetAllErrorCodeResp struct {
	Code map[int]string `json:"code"`
}
