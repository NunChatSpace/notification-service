package deliveries_noti

import (
	"net/http"

	"github.com/NunChatSpace/notification-service/internal/entities"
	"github.com/NunChatSpace/notification-service/internal/response_wrapper"
)

func DestinationList(db entities.DB) (response_wrapper.Model, error) {
	return response_wrapper.Model{
		StatusCode: http.StatusOK,
		Data:       nil,
		Message:    "",
	}, nil
}

func SourceList(db entities.DB) (response_wrapper.Model, error) {
	return response_wrapper.Model{
		StatusCode: http.StatusOK,
		Data:       nil,
		Message:    "",
	}, nil
}

func AddDestination(db entities.DB, model Destination) (response_wrapper.Model, error) {

	return response_wrapper.Model{
		StatusCode: http.StatusOK,
		Data:       model,
		Message:    "",
	}, nil
}

func AddSource(db entities.DB, model Source) (response_wrapper.Model, error) {

	return response_wrapper.Model{
		StatusCode: http.StatusOK,
		Data:       model,
		Message:    "",
	}, nil
}

func AddExpression(db entities.DB, model Expression) (response_wrapper.Model, error) {

	return response_wrapper.Model{
		StatusCode: http.StatusOK,
		Data:       model,
		Message:    "",
	}, nil
}
