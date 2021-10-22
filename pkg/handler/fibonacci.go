package handler

import (
	"net/http"

	"github.com/elli56/fibo-api/pkg/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) fibonacciCalc(c *gin.Context) {
	// result := make(map[int64]int64)
	var input entities.FibonacciInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	result, err := h.services.Calculation.FiboSlice(input.X, input.Y)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"input":  input,
		"result": result,
	})
}

func (h *Handler) test(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
