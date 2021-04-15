package handler

import (
	"battleship/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ships struct {
	Coordinates string `json:"Coordinates"`
}

func (h *Handler) addShipInBoard(c *gin.Context) {
	var coordinateShips Ships

	if err := c.BindJSON(&coordinateShips); err != nil {
		logger.Error("invalid input body")
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid input body")
	}

	cs := coordinateShips.Coordinates
	err := h.services.AddShipInBoard(cs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "false"})
		return
	}
	err2 := h.services.ShowBoard()
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "false"})
		return
	}
}
