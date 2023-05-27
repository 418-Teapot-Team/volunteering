package app

import (
	"fmt"
	"github.com/BoryslavGlov/logrusx"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	logg logrusx.Logging
}

func NewApp(logg logrusx.Logging) *Handler {
	return &Handler{logg: logg}
}

func (app *App) newErrorResponse(c *gin.Context, statusCode int, message string) {
	app.logg.Error(message,
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
