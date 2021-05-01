package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/aws/aws-lambda-go/lambdacontext"
	log "github.com/sirupsen/logrus"
	"keyforge-cards-backend/internal/api"
	"keyforge-cards-backend/internal/service"
	"net/http"
	"strings"
)

func handleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	parameters := event.QueryStringParameters
	ttsObject, err := service.ImportDeck(parameters["deckid"], language(event), defaultIfBlank(parameters["sleeve"], "red"))

	if err != nil {
		log.Error(err.Error())
		return api.Error(api.StatusCodeFromError(err), nil, "failed to fetch deck", err), err
	}

	return api.Response(http.StatusOK, nil, ttsObject), nil
}

func language(event events.APIGatewayProxyRequest) string {
	if val, ok := event.Headers["Accept-Language"]; ok {
		return val
	} else {
		return "pt"
	}
}

func defaultIfBlank(value string, defaultValue string) string {
	trimmedValue := strings.TrimSpace(value)
	if trimmedValue == "" {
		return defaultValue
	} else {
		return trimmedValue
	}
}

func main() {
	lambda.Start(handleRequest)
}
