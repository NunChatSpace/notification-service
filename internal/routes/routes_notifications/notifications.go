package routes_notifications

import (
	"github.com/gin-gonic/gin"
)

func AddGroup(r *gin.Engine) {
	notificationGroup := r.Group("/notifications")

	notificationGroup.GET("/source/:uid", getSource)
	notificationGroup.GET("/source", sourceList)
	notificationGroup.POST("/source", addSource)

	notificationGroup.GET("/destination/:uid", getDestination)
	notificationGroup.GET("/destination", destinationList)
	notificationGroup.POST("/destination", addDestination)

	notificationGroup.GET("/expression/:uid", getExpression)
	notificationGroup.GET("/expression", getAllExpression)
	notificationGroup.POST("/expression", addExpression)
}
