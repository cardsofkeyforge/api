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

func Error(statusCode int, headers map[string]string, message string, err error) events.APIGatewayProxyResponse {
	return Response(statusCode, headers, fmt.Sprintf("{ %q : %q, %q : %q }", "message", message, "error", err.Error()))
}

func StatusCodeFromError(err error) int {
	switch errors2.Cause(err).(type) {
	case *dynamodb.ResourceNotFoundException:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
