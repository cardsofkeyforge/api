package api

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	errors2 "github.com/pkg/errors"
	"net/http"
)

func Response(statusCode int, headers map[string]string, body interface{}) events.APIGatewayProxyResponse {
	marshal, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       fmt.Sprintf("%s", marshal),
	}
}

func Error(statusCode int, message string, headers map[string]string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       fmt.Sprintf("{ %q : %s }", "message", message),
	}
}

func StatusCodeFromError(err error) int {
	switch errors2.Cause(err).(type) {
	case *dynamodb.ResourceNotFoundException:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}