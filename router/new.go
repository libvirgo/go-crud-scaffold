package router

import (
	"context"
	"fmt"
	"frame/conf"
	"frame/docs"
	"frame/router/middleware"
	"frame/router/user"
	validator2 "frame/router/validator"
	"frame/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"time"
)

// New is a fx option to provide router module
// @title Frame API
// @version v1
// @host localhost:4000
// @BasePath  /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func New() fx.Option {
	var module = fx.Module("server",
		fx.Provide(
			middleware.NewAuthMiddleware,
			middleware.NewJwtClaimsMiddleware,
		),
		// routers
		fx.Provide(
			utils.AsRouter(user.NewUserRouter),
		),
		fx.Provide(
			fx.Annotate(
				NewRouter,
				fx.ParamTags(`group:"routers"`),
			)),
		fx.Provide(
			NewEngine,
			NewApiGroup,
			utils.NewHandler,
		),
		fx.Provide(
			NewHTTPServer,
		),
		fx.Invoke(func(router *http.Server) {}),
	)
	return module
}

func NewEngine(logger *zap.Logger, conf *conf.Config, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	if conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("wallet", validator2.WalletValidator)
		if err != nil {
			logger.Error("register wallet validator failed", zap.Error(err))
		}
	}
	// ginzap middleware for gin
	r.Use(ginzap.Ginzap(logger, time.RFC3339, conf.Debug))
	r.Use(ginzap.RecoveryWithZap(logger, false))
	r.Use(middleware.NewCors())

	if conf.Debug {
		docs.SwaggerInfo.BasePath = "/api/v1"
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(context *gin.Context) {
		claims := jwt.ExtractClaims(context)
		logger.Info("NoRoute claims", zap.Any("claims", claims))
		context.JSON(404, gin.H{
			"code": 404,
			"data": "Page not found",
		})
	})
	return r
}

func NewApiGroup(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	r.Use(middleware.ResponseMiddleware())
	g := r.Group("/api/v1")
	g.GET("/errorcode", authMiddleware.MiddlewareFunc(), GetAllErrorCode)
	return r.Group("/api/v1")
}

// NewRouter ensure all router registered
func NewRouter(_ []utils.Router) utils.Router {
	return utils.Void{}
}

// NewHTTPServer create http server, the last argument is a void to ensure NewRouter invoked
func NewHTTPServer(lc fx.Lifecycle, router *gin.Engine, conf *conf.Config, _ utils.Router) *http.Server {
	srv := &http.Server{Addr: fmt.Sprintf(":%d", conf.Port), Handler: router}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				fmt.Println("Failed to start HTTP Server at:", srv.Addr)
				return err
			}
			go func() {
				err := srv.Serve(ln)
				if err != nil {
					log.Println("Failed to start HTTP Server at:", srv.Addr, "error:", err)
				} else {
					log.Println("Succeed to start HTTP Server at:", srv.Addr)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := srv.Shutdown(ctx)
			if err != nil {
				log.Println("Failed to stop HTTP Server at:", srv.Addr, "error:", err)
				return err
			}
			log.Println("Succeed to stop HTTP Server at:", srv.Addr)
			return nil
		},
	})
	return srv
}
