package routes_notifications

import (
	"github.com/gin-gonic/gin"
)

func AddGroup(r *gin.Engine) {
	notificationGroup := r.Group("/notifications")

	notificationGroup.GET("/source", sourceList)
	notificationGroup.POST("/source", addSource)

	notificationGroup.GET("/destination", destinationList)
	notificationGroup.POST("/destination", addDestination)

	notificationGroup.POST("/expression", addExpression)
}

func sourceList(c *gin.Context) {

}

func addSource(c *gin.Context) {

}

func destinationList(c *gin.Context) {

}

func addDestination(c *gin.Context) {

}

func addExpression(c *gin.Context) {

}
