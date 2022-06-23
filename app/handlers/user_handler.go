// Package handlers User-service Api:
//
//   version: 0.0.1
//   title: User-service Api
//  Schemes: http, https
//  Host: localhost:8080
//  BasePath: /
//  Produces:
//    - application/json
//	Consumes:
//    - application/json
//
// securityDefinitions:
//  Bearer:
//    type: apiKey
//    in: header
//    name: authorization
//
// swagger:meta
package handlers

import (
	"net/http"
	"strings"

	pb "github.com/badfan/inno-taxi-user-service/app/rpc"

	"github.com/pkg/errors"

	"github.com/badfan/inno-taxi-user-service/app/models"
	"github.com/gin-gonic/gin"
)

// swagger:route POST /sign-up/ user SignUp
//
// Sign up
//
// responses:
//   200: idResponse
//   400: ErrorMsg
//   500: ErrorMsg
func (h *Handler) SignUp(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorLogger(c, err)
		return
	}

	id, err := h.userService.SignUp(c.Request.Context(), &input)
	if err != nil {
		h.ErrorLogger(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// swagger:route POST /sign-in/ user SignIn
//
// Sign in
//
// responses:
//   200: tokenResponse
//   404: ErrorMsg
//   500: ErrorMsg
func (h *Handler) SignIn(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorLogger(c, err)
		return
	}

	token, err := h.userService.SignIn(c.Request.Context(), input.PhoneNumber, input.Password)
	if err != nil {
		h.ErrorLogger(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// swagger:route GET /user/api/v1/rating/ user GetUserRating
//
// Getting user rating
//
// security:
//  - Bearer:
// 		- []
//
// responses:
//   200: ratingResponse
//   500: ErrorMsg
func (h *Handler) GetUserRating(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		h.ErrorLogger(c, errors.New("user id not found"))
		return
	}

	rating, err := h.userService.GetUserRating(c.Request.Context(), id.(int))
	if err != nil {
		h.ErrorLogger(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"rating": rating,
	})
}

// swagger:route POST /user/api/v1/set-rating/ user SetDriverRating
//
// Setting driver rating
//
// security:
//  - Bearer:
//		- []
//
// responses:
//   200: msgResponse
//   500: ErrorMsg
func (h *Handler) SetDriverRating(c *gin.Context) {
	_, ok := c.Get("userID")
	if !ok {
		h.ErrorLogger(c, errors.New("user id not found"))
		return
	}

	var input pb.SetDriverRatingRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorLogger(c, err)
		return
	}

	err := h.userService.SetDriverRating(c.Request.Context(), int(input.GetRating()))
	if err != nil {
		h.ErrorLogger(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "your rating is accepted",
	})
}

// swagger:route GET /user/api/v1/order-history/ user GetOrderHistory
//
// Getting order history
//
// security:
//  - Bearer:
// 		- []
//
// responses:
//   200: ordersResponse
//   500: ErrorMsg
func (h *Handler) GetOrderHistory(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		h.ErrorLogger(c, errors.New("user id not found"))
		return
	}

	orderHistory, err := h.userService.GetOrderHistory(c.Request.Context(), id.(int))
	if err != nil {
		h.ErrorLogger(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"orders": strings.Join(orderHistory, "\n"),
	})
}

// swagger:route POST /user/api/v1/find-taxi/ user FindTaxi
//
// Finding taxi
//
// security:
//  - Bearer:
// 		- []
//
// responses:
//   200: findTaxiResponse
//   500: ErrorMsg
func (h *Handler) FindTaxi(c *gin.Context) {
	id, ok := c.Get("userID")
	if !ok {
		h.ErrorLogger(c, errors.New("user id not found"))
		return
	}

	var input pb.GetTaxiForUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.ErrorLogger(c, err)
		return
	}

	driverUUID, driverRating, err := h.userService.FindTaxi(c.Request.Context(), id.(int), input.GetOrigin(), input.GetDestination(), input.GetTaxiType())
	if err != nil {
		h.ErrorLogger(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"driver_uuid": driverUUID,
		"rating":      driverRating,
	})
}

// swagger:response idResponse
type _ struct {
	// in: body
	Body struct {
		ID int `json:"id,omitempty"`
	}
}

// swagger:response tokenResponse
type _ struct {
	// in: body
	Body struct {
		Token string `json:"token,omitempty"`
	}
}

// swagger:response ratingResponse
type _ struct {
	// in: body
	Body struct {
		Rating float32 `json:"rating,omitempty"`
	}
}

// swagger:response msgResponse
type _ struct {
	// in: body
	Body struct {
		Message string `json:"message,omitempty"`
	}
}

// swagger:response ordersResponse
type _ struct {
	// in: body
	Body struct {
		Orders []string `json:"orders,omitempty"`
	}
}

// swagger:response findTaxiResponse
type _ struct {
	// in: body
	Body struct {
		DriverUUID string  `json:"driver_uuid,omitempty"`
		Rating     float32 `json:"rating,omitempty"`
	}
}

// swagger:response ErrorMsg
type _ struct {
	// in: body
	Body struct {
		Message string `json:"message,omitempty"`
	}
}

// swagger:parameters SignUp
type _ struct {
	// in: body
	Body struct {
		Name        string `json:"name,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		Email       string `json:"email,omitempty"`
		Password    string `json:"password,omitempty"`
	}
}

// swagger:parameters SignIn
type _ struct {
	// in: body
	Body struct {
		PhoneNumber string `json:"phone_number,omitempty"`
		Password    string `json:"password,omitempty"`
	}
}

// swagger:parameters SetDriverRating
type _ struct {
	// in: body
	Body struct {
		Rating int `json:"driver_rating,omitempty"`
	}
}

// swagger:parameters FindTaxi
type _ struct {
	// in: body
	Body struct {
		Origin      string `json:"origin,omitempty"`
		Destination string `json:"destination,omitempty"`
		TaxiType    string `json:"taxi_type,omitempty"`
	}
}
