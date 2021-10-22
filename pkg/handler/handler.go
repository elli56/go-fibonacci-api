package handler

import (
	"github.com/elli56/fibo-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/test", h.test)

	api := router.Group("/api")
	{
		// api.GET("/", h.Index)
		api.POST("/fibonacci-calc", h.fibonacciCalc)
	}
	// возвращает type pointer  gin.Engine
	return router
}
