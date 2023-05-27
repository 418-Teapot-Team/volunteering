package app

import (
	"fmt"
	"github.com/BoryslavGlov/logrusx"
	"github.com/gin-gonic/gin"
	"sync"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {

	var mu sync.Mutex

	mu.Lock()
	h.logg.Error(message,
		logrusx.LogField{
			Key:   "statusCode",
			Value: statusCode},
		logrusx.LogField{
			Key:   "request",
			Value: fmt.Sprintf("%+v", c.Request),
		},
	)
	mu.Unlock()
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
