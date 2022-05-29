package v1

import "github.com/gin-gonic/gin"

func (a *ApiV1) UserGroup(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	{
		userGroup.GET("/rating", a.handler.GetUserRating)
	}
}
