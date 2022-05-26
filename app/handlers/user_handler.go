package handlers

import (
	"net/http"

	"github.com/badfan/inno-taxi-user-service/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.SignUp(c.Request.Context(), &input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.SignIn(c.Request.Context(), input.PhoneNumber, input.Password)
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

	rating, err := h.service.GetUserRating(c.Request.Context(), conID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"rating": rating,
	})
}
