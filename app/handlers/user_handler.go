package handlers

import (
	"github.com/badfan/inno-taxi-user-service/app/models/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input sqlc.User

	if err := c.ShouldBindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.SignUp(input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input sqlc.User

	if err := c.ShouldBindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.SignIn(input.PhoneNumber, input.Password)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) GetUserRating(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	conID, ok := id.(int)
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "cannot convert id to type int")
		return
	}

	rating, err := h.service.GetUserRating(conID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"rating": rating,
	})
}
