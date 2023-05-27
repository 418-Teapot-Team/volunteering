package app

import (
	"fmt"
	"github.com/BoryslavGlov/logrusx"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {

	h.logg.Error(message,
		logrusx.LogField{
			Key:   "statusCode",
			Value: statusCode},
		logrusx.LogField{
			Key:   "request",
			Value: fmt.Sprintf("%+v", c.Request),
		},
	)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
