package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rsaraphael/cardap/db"
	"github.com/rsaraphael/cardap/model"
)

type GetUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func HandleRequest(ctx context.Context, request GetUserRequest) ([]model.User, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	defer db2.Close()
	if err != nil {
		return []model.User{}, err
	}
	var users []model.User
	db2.Find(&users)
	return users, nil
}
func main() {
	lambda.Start(HandleRequest)
}
