package handlers

import (
	"github.com/NunChatSpace/notification-service/http/handlers/cors"
	"github.com/NunChatSpace/notification-service/http/handlers/database"
	"github.com/NunChatSpace/notification-service/internal/entities"
	"github.com/gin-gonic/gin"
)

func SetHandlers(db entities.DB) *gin.Engine {
	router := gin.Default()

	router.Static("/static", "./static")
	router.Use(cors.Handler())
	router.Use(database.Handler(db))

	return router
}
