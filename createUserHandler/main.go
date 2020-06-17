package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rsaraphael/cardap/db"
	"github.com/rsaraphael/cardap/model"
)

func HandleRequest(ctx context.Context) (model.User, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	defer db2.Close()
	if err != nil {
		return model.User{}, err
	}
	account := &model.User{Email: "raphael@email.com", Pass: "123", Username: "rsaraphael"}
	db2.Create(account)
	return *account, nil
}

func main() {
	lambda.Start(HandleRequest)
}
