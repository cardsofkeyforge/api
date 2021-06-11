package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/aws/aws-lambda-go/lambdacontext"
	log "github.com/sirupsen/logrus"
	"keyforge-cards-backend/internal/api"
	"keyforge-cards-backend/internal/service"
	"net/http"
)

func handleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cards, err := service.SearchCards(event)

	headers := map[string]string{
		"Access-Control-Allow-Origin":  "https://site.cardsofkeyforge.com",
		"Access-Control-Allow-Methods": "*",
		"Access-Control-Allow-Headers": "*",
	}

	if err != nil {
		log.Error(err.Error())
		return api.Error(api.StatusCodeFromError(err), headers, "failed to fetch cards", err), err
	}

	return api.Response(http.StatusOK, headers, cards), nil
}

func main() {
	lambda.Start(handleRequest)
}
