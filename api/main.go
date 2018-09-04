package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dougireton/sentences/sentences"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	log.Printf("Body size = %d.\n", len(request.Body))

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%v: Missing body text", http.StatusText(http.StatusBadRequest)),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	s := sentences.ParseText(request.Body)
	sentences, err := json.Marshal(s)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(sentences),
		StatusCode: http.StatusOK,
	}, nil

}

func main() {
	lambda.Start(handler)
}
