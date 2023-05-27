package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	//token, err := h.s.GenerateToken(input.Email, input.Password)
	//if err != nil {
	//	h.newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
