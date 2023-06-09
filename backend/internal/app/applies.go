package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"volunteering"
)

func (h *Handler) applyToTask(c *gin.Context) {
	var input volunteering.Applies

	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.AppliedUserID = userId
	input.CreatedAt = time.Now()

	err = h.rep.MakeApply(&input)
	if err != nil {
		h.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}

func (h *Handler) getApplies(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	list, err := h.rep.GetAllApplies(userId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": list,
	})

}

func (h *Handler) approveApply(c *gin.Context) {

	var input struct {
		Id        int `json:"id"`
		AppliedId int `json:"applied_id"`
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

	if input.AppliedId == userId {
		h.newErrorResponse(c, http.StatusForbidden, "You can't apply yourself")
		return
	}

	err = h.rep.ApproveApply(userId, input.Id, input.AppliedId)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}
