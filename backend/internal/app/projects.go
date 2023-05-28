package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"volunteering"
)

func (h *Handler) createProject(c *gin.Context) {
	var input volunteering.Project

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

	err = h.rep.CreateProject(&input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (h *Handler) deleteProject(c *gin.Context) {
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

	err = h.rep.DeleteProject(input.Id, userId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}

func (h *Handler) getProjects(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	projects, err := h.rep.GetTasks(userId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": projects,
	})

}

func (h *Handler) deleteProjectTask(c *gin.Context) {
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

	err = h.rep.DeleteTaskProject(userId, input.Id)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": "Success",
	})
}

func (h *Handler) getProjectStats(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	projects, err := h.rep.GetProjectStats(userId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": projects,
	})

}
