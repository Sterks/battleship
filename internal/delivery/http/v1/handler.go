package v1

import (
	"battleship/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	router.POST("/create-matrix", h.createMatrix)
}

func (h *Handler) createMatrix(c *gin.Context) {
	var rangeInt int
	if err := c.BindJSON(&rangeInt); err != nil {
		logger.Error("invalid input body")
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid input body" )
	}

	h.createMatrix()

}
