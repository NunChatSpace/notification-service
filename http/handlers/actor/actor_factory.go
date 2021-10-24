package actor

import (
	"github.com/NunChatSpace/notification-service/config"
	"github.com/NunChatSpace/notification-service/internal/jwt_token"
)

type Model struct {
	UserID     string
	Permission []string
}

func GetActorFromToken(token string, config *config.Config) Model {
	claims, err := jwt_token.Decode(token, config)
	if err != nil {
		return Model{}
	}

	return Model{
		UserID:     claims.UserID,
		Permission: claims.Permission,
	}
}
