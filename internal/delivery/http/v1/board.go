package handler

import (
	"battleship/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	if err2 := h.services.CreateBoard(rI); err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "false"})
		return
	}
	err := h.services.ShowBoard()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "false"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
