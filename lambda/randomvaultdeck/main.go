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

var setName = map[string]int{
	"all":     0,
	"cota":    341,
	"aoa":     435,
	"wc":      452,
	"anomaly": 453,
	"mm":      479,
	"dt":      496,
	"rotk":    1001,
}

func handleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	parameters := event.QueryStringParameters
	deckId, err := service.RetrieveRandomDeckId(defaultIfBlank(parameters["set"], "all"))

	if err != nil {
		log.Error(err.Error())
		return api.Error(api.StatusCodeFromError(err), nil, "failed to fetch deck id", err), err
	}

	return api.Response(http.StatusOK, nil, deckId), nil
}

func defaultIfBlank(value string, defaultValue string) int {
	trimmedValue := strings.TrimSpace(value)
	if trimmedValue == "" {
		return setName[defaultValue]
	} else {
		return setName[trimmedValue]
	}
}

func main() {
	lambda.Start(handleRequest)
}
