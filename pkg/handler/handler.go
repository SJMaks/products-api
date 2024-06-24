package handler

import (
	"github.com/SJMaks/products-api/pkg/service"
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

	api := router.Group("/api")
	{
		product := api.Group("/product")
		{
			product.GET("/", h.getProducts)
			product.GET("/:i", h.getProduct)
			product.POST("/", h.createProduct)
			product.PUT("/:i", h.updateProduct)
			product.DELETE("/:i", h.deleteProduct)
		}
		measure := api.Group("/measure")
		{
			measure.GET("/", h.getMeasures)
			measure.GET("/:i", h.getMeasure)
			measure.POST("/", h.createMeasure)
			measure.PUT("/:i", h.updateMeasure)
			measure.DELETE("/:i", h.deleteMeasure)
		}
	}

	return router
}
