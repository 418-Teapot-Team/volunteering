package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"volunteering"
)

func (h *Handler) createTask(c *gin.Context) {
	var input volunteering.Task

	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserId = userId
	input.CreatedAt = time.Now()

	err = h.rep.CreateTask(&input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (h *Handler) deleteTask(c *gin.Context) {
	var input struct {
		Id int `json:"id"`
	}

	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.rep.DeleteTask(input.Id, userId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}
