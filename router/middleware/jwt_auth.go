package middleware

import (
	"errors"
	"fmt"
	"frame/conf"
	"frame/ent"
	"frame/ent/adminuser"
	"frame/ent/user"
	"frame/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

func NewAuthMiddleware(logger *zap.Logger, conf *conf.Config, c *ent.Client, rd *redis.Client, q *asynq.Client) *jwt.GinJWTMiddleware {
	authMiddleWare, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "frame",
		Key:         []byte(conf.JWTSecret),
		Timeout:     time.Hour * time.Duration(conf.JWTExpireHour),
		IdentityKey: utils.GinPayloadCtxKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*utils.GinPayload); ok {
				return jwt.MapClaims{
					utils.GinWalletAddressCtxKey: v.WalletAddress,
					utils.GinIsAdminCtxKey:       v.IsAdmin,
					utils.GinUserIdCtxKey:        v.UserId,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(context *gin.Context) interface{} {
			claims := jwt.ExtractClaims(context)
			return &utils.GinPayload{
				WalletAddress: claims[utils.GinWalletAddressCtxKey].(string),
				IsAdmin:       claims[utils.GinIsAdminCtxKey].(bool),
				UserId:        int(claims[utils.GinUserIdCtxKey].(float64)),
			}
		},
		Authenticator: func(context *gin.Context) (interface{}, error) {
			var loginVal LoginReq
			if err := context.ShouldBindJSON(&loginVal); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			// debug mode not to verify sig.
			if !conf.Debug {
				nonceKey := rd.Get(context, fmt.Sprintf(utils.WalletLoginNonceKey, loginVal.WalletAddress)).Val()
				if nonceKey == "" {
					return "", errors.New("nonce not found")
				}
				defer func() {
					rd.Del(context, fmt.Sprintf(utils.WalletNonceKey, loginVal.WalletAddress))
				}()
				if !utils.VerifySig(loginVal.WalletAddress, loginVal.SignData, []byte(nonceKey)) {
					return "", errors.New("verify sig failed")
				}
			}

			u, err := c.User.Query().Where(
				user.WalletAddressEQ(loginVal.WalletAddress),
			).First(context)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			go func() {
				_, err := c.UserActivity.Create().SetUser(u).SetType(utils.UserActivityLogin).Save(context)
				if err != nil {
					logger.Error("create user activity error", zap.String("wallet_address", u.WalletAddress), zap.Error(err))
				}
			}()
			var payload = utils.GinPayload{
				WalletAddress: u.WalletAddress,
				UserId:        u.ID,
			}
			_, err = c.AdminUser.Query().Where(
				adminuser.WalletAddressEQ(u.WalletAddress)).First(context)
			if err == nil {
				payload.IsAdmin = true
			}
			return &payload, nil
		},
		Authorizator: func(data interface{}, context *gin.Context) bool {
			if _, ok := data.(*utils.GinPayload); ok {
				return true
			}
			return false
		},
		Unauthorized: func(context *gin.Context, code int, message string) {
			context.JSON(code, message)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		panic(err)
	}
	return authMiddleWare
}
