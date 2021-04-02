package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	_ "github.com/aws/aws-lambda-go/lambdacontext"
)

type Card struct {
	Name   string `json:"name"`
	Set    string `json:"set"`
	Number string `json:"number"`
}

func handleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	parameters := event.PathParameters
	id := parameters["id"]
	set := parameters["set"]
	bytes, _ := json.Marshal(Card{
		Name:   "Cardzinho",
		Set:    "Setzinho",
		Number: id,
	})

	if set == "cota" {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       fmt.Sprintf("%s", bytes),
		}, nil
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
		}, nil
	}
}

func main() {
	runtime.Start(handleRequest)
}
