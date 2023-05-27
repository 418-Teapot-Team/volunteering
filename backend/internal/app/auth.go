package app

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
	"volunteering"
	"volunteering/pkg/repository"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var (
		input volunteering.User
		mu    sync.Mutex
	)

	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.Password = generatePasswordHash(input.Password)
	input.CreatedAt = time.Now()
	input.LastLogin = time.Now()
	mu.Unlock()

	err := h.rep.CreateUser(&input)
	if err != nil {
		if errors.Is(err, repository.UniqueViolationError) {
			h.newErrorResponse(c, http.StatusConflict, "user is already exists")
			return
		}
		h.newErrorResponse(c, http.StatusInternalServerError, "something gone wrong while trying create user")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.GenerateToken(input.Email, input.Password)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) whoAmI(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		h.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	user, err := h.rep.GetUserById(userId)
	if err != nil || user.Id == 0 {
		h.newErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"userId":    user.Id,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
	})
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
