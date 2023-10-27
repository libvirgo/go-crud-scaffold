package middleware

import (
	"encoding/json"
	"frame/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

func NewJwtClaimsMiddleware(j *jwt.GinJWTMiddleware) JwtClaimsMiddleware {
	return func(context *gin.Context) {
		claims, err := j.GetClaimsFromJWT(context)
		if err != nil {
			return
		}
		// if claims expired, return
		switch v := claims["exp"].(type) {
		case nil:
			return
		case float64:
			if int64(v) < time.Now().Unix() {
				return
			}
		case json.Number:
			n, err := v.Int64()
			if err != nil {
				return
			}
			if n < time.Now().Unix() {
				return
			}
		default:
			return
		}
		payload := utils.GinPayload{
			WalletAddress: claims[utils.GinWalletAddressCtxKey].(string),
			IsAdmin:       claims[utils.GinIsAdminCtxKey].(bool),
			UserId:        int(claims[utils.GinUserIdCtxKey].(float64)),
		}
		context.Set(utils.GinPayloadCtxKey, &payload)
	}
}
