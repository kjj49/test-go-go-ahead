package http

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error string `json:"error" example:"message"`
}

func newErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, errorResponse{msg})
}
