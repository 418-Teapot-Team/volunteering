package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		h.newErrorResponse(c, http.StatusForbidden, "empty auth header")
		return
	}
	headersParts := strings.Split(header, " ")
	if len(headersParts) != 2 || headersParts[0] != "Bearer" {
		h.newErrorResponse(c, http.StatusForbidden, "invalid auth header")
		return
	}

	if len(headersParts[1]) == 0 {
		h.newErrorResponse(c, http.StatusForbidden, "token is empty")
		return
	}

	userId, err := h.ParseToken(headersParts[1])
	if err != nil {
		h.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.Set(userCtx, userId.UserId)
}

func (h *Handler) getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	integer, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return integer, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
