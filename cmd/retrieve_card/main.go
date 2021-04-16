package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/aws/aws-lambda-go/lambdacontext"
	log "github.com/sirupsen/logrus"
	"keyforge-cards-backend/internal/api"
	"keyforge-cards-backend/internal/database"
	"keyforge-cards-backend/internal/models"
	"net/http"
	"strings"
)

func handleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var results []models.Card
	var tableName string
	fb := database.FilterBuilder{}


	name := event.PathParameters["cardName"]
	set := event.PathParameters["set"]

	if val, ok := event.Headers["Lang"]; ok {
		tableName = fmt.Sprintf("cards_%s", val)
	} else {
		tableName = fmt.Sprintf("cards_%s", "pt")
	}

	expression, values, err := fb.Contains("CardTitle", strings.Title(strings.ToLower(name))).
		And().
		Eq("Set", strings.ToLower(set)).Build()

	if err != nil {
		log.Error(err.Error())
		return api.Error(api.StatusCodeFromError(err), err.Error(), nil), err
	}

	err = database.Scan(tableName, expression, values, &results)

	if err != nil {
		log.Error(err.Error())
		return api.Error(api.StatusCodeFromError(err), err.Error(), nil), err
	}

	return api.Response(http.StatusOK, nil, results), nil

}

func main() {
	lambda.Start(handleRequest)
}
