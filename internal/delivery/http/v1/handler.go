package handler

import (
	"battleship/internal/service"
	"battleship/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
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
	return router
}

type req struct {
	RangeInt int `json:"range"`
}

func (h *Handler) createMatrix(c *gin.Context) {
	var rangeInt req
	if err := c.BindJSON(&rangeInt); err != nil {
		logger.Error("invalid input body")
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid input body")
	}

	rI := rangeInt.RangeInt

	if err2 := h.services.ICreateMatrix.CreateMatrix(rI); err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "false"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
