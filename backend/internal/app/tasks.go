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

func (h *Handler) getUserTasks(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tasks, err := h.rep.GetUserTasks(userId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": tasks,
	})

}

func (h *Handler) markAsDoneVolunteer(c *gin.Context) {
	var input struct {
		Id          int `json:"id"`
		LoggedHours int `json:"loggedHours"`
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

	err = h.rep.MarkAsDoneVolunteer(userId, input.Id, input.LoggedHours)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}

func (h *Handler) markAsDoneEmployer(c *gin.Context) {
	var input struct {
		Id          int `json:"id"`
		LoggedHours int `json:"loggedHours,omitempty"`
	}
	var d bool
	done := c.Query("done")
	if done == "" || done == "false" {
		d = false
	} else if done == "true" {
		d = true
	} else {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid parameter done")
		return
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

	err = h.rep.MarkAsDoneEmployer(userId, input.Id, input.LoggedHours, d)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}

func (h *Handler) shareTask(c *gin.Context) {
	var input struct {
		Id    int  `json:"id"`
		Share bool `json:"share"`
	}

	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.rep.ShareTask(input.Id, input.Share, userId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (h *Handler) getSharedTasks(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tasks, err := h.rep.GetSharedTasks(userId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": tasks,
	})

}
