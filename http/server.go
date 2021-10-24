package http

import (
	"github.com/NunChatSpace/notification-service/http/handlers"
	"github.com/NunChatSpace/notification-service/internal/entities"
	"github.com/NunChatSpace/notification-service/internal/routes/routes_notifications"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	db, err := entities.NewDB()
	if err != nil {
		panic(err)
	}

	router := handlers.SetHandlers(db)
	routes_notifications.AddGroup(router)

	router.GET("/_healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthy",
		})
	})

	return router
}
