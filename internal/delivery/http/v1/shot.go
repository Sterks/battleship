package handler

import (
	"battleship/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Shot struct {
	Coord string `json:"Coord"`
}

func (h *Handler) shot(c *gin.Context) {
	var coord Shot
	if err := c.BindJSON(&coord); err != nil {
		logger.Error("invalid input body")
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid input body")
	}

	coo := coord.Coord

	err := h.services.ICreateBoard.Shot(coo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "false"})
		return
	}

	hhs := h.services.ICreateBoard.ShotResult(coo)

	c.JSON(http.StatusOK, hhs)

	err2 := h.services.ICreateBoard.ShowBoard()
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "false"})
		return
	}
}
