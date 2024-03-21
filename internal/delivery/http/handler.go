package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/paw1a/ecommerce-api/docs"
	"github.com/paw1a/ecommerce-api/internal/config"
	v1 "github.com/paw1a/ecommerce-api/internal/delivery/http/v1"
	"github.com/paw1a/ecommerce-api/internal/service"
	"github.com/paw1a/ecommerce-api/pkg/auth"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"
)

type Handler struct {
	config        *config.Config
	services      *service.Services
	tokenProvider auth.TokenProvider
}

func NewHandler(services *service.Services,
	tokenProvider auth.TokenProvider, config *config.Config) *Handler {

	return &Handler{
		config:        config,
		services:      services,
		tokenProvider: tokenProvider,
	}
}

func (h *Handler) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhsot:3000"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenProvider, h.config)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
