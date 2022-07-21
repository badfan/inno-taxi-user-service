package v1

import "github.com/gin-gonic/gin"

func (a *ApiV1) UserGroup(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	{
		userGroup.GET("/order-history", a.handler.GetOrderHistory)
		userGroup.GET("/rating", a.handler.GetUserRating)
		userGroup.POST("/find-taxi", a.handler.FindTaxi)
		userGroup.POST("/set-rating", a.handler.SetDriverRating)
	}
}
