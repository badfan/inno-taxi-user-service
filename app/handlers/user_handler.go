package handlers

import (
	"net/http"

	pb "github.com/badfan/inno-taxi-user-service/app/services/proto"

	"github.com/badfan/inno-taxi-user-service/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {

		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.userService.SignUp(c.Request.Context(), &input)
	if err != nil {
		h.logger.Errorf("error occured while signing up: %s", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "unable to sign up")
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

	token, err := h.userService.SignIn(c.Request.Context(), input.PhoneNumber, input.Password)
	if err != nil {
		h.logger.Errorf("error occured while signing in: %s", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "unable to sign in")
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

	rating, err := h.userService.GetUserRating(c.Request.Context(), id.(int))
	if err != nil {
		h.logger.Errorf("error occured while getting user rating: %s", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "unable to get user rating")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"rating": rating,
	})
}

func (h *Handler) SetDriverRating(c *gin.Context) {
	_, ok := c.Get("userID")
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input pb.SetDriverRatingRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.userService.SetDriverRating(c.Request.Context(), int(input.GetRating()))
	if err != nil {
		h.logger.Errorf("error occured while setting driver's rating: %s", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "unable to set driver's rating")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "your rating is accepted",
	})
}

func (h *Handler) GetOrderHistory(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	orderHistory, err := h.userService.GetOrderHistory(c.Request.Context(), id.(int))
	if err != nil {
		h.logger.Errorf("error occured while getting order history: %s", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "unable to get order history")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"orders": orderHistory,
	})
}

func (h *Handler) GetTaxi(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input pb.GetTaxiForUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	info, err := h.userService.GetTaxi(c.Request.Context(), id.(int), input.GetOrigin(), input.GetDestination(), input.GetTaxiType())
	if err != nil {
		h.logger.Errorf("error occured while getting taxi: %s", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "unable to get taxi")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"driver_uuid": info.GetDriverUuid(),
		"rating":      info.GetDriverRating(),
	})
}
