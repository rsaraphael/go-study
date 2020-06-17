package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rsaraphael/cardap/db"
	"github.com/rsaraphael/cardap/model"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	resp := events.APIGatewayProxyResponse{}
	defer db2.Close()
	if err != nil {
		return resp, err
	}
	account := model.User{Email: "raphael@email.com", Pass: "123", Username: req.PathParameters["name"]}
	db2.Create(&account)

	fmt.Print("%v", req.PathParameters)
	body, err := json.Marshal(&account)
	resp.Body = string(body)
	resp.StatusCode = 200
	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
