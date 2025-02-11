package app

import (
	"time"

	"github.com/Crabocod/golang-test/config"
	"github.com/Crabocod/golang-test/internal/cache"
	"github.com/Crabocod/golang-test/internal/handler"
	"github.com/Crabocod/golang-test/internal/service"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	config *config.Config
}

func New(config *config.Config) *App {
	return &App{
		router: gin.Default(),
		config: config,
	}
}

func (a *App) Run() error {
	cacheService := cache.NewCacheService(
		a.config.Redis.Host,
		a.config.Redis.Port,
		a.config.Redis.DB,
		24*time.Hour,
	)

	hashService := service.NewHashService(cacheService)
	hashHandler := handler.NewHashHandler(hashService)

	a.router.POST("/hash", hashHandler.CreateHash)

	return a.router.Run(a.config.Server.Port)
}
