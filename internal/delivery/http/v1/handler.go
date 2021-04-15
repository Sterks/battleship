package handler

import (
	"battleship/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	router.POST("/create-matrix", h.createMatrix)
	router.POST("/ship", h.addShipInBoard)
	router.POST("/shot", h.shot)
	router.POST("/clear", h.clear)
	return router
}
