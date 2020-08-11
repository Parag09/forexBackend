package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// forex data struct
type ForexData struct {
	Curr      string `json:"Curr"`
	Rate      string `json:"Rate"`
	Timestamp string `json:"Timestamp"`
}

type ForexDataInput struct {
	Rates map[string]interface{} `json:"rates"`
}

type currObj struct {
	Rate string `json:"rate"`
	Time string `json:"timestamp"`
}

// function to get single forex item based on query param
func getForexItem() (*ForexData, error) {

	forexitem, err := getItem("EURGBP")
	if err != nil {
		return nil, err
	}
	return forexitem, nil
}

// function to get all the forex items
func getForexAllItem() (events.APIGatewayProxyResponse, error) {

	// Fetch the forex record from the database.
	forexItems, err := getAllItem()
	if err != nil {
		return serverError(err)
	}
	if forexItems == nil {
		return clientError(http.StatusNotFound)
	}

	// The APIGatewayProxyResponse.Body field needs to be a string, so
	// we marshal the forex record into JSON.
	js, err := json.Marshal(forexItems)
	fmt.Println(js)
	if err != nil {
		return serverError(err)
	}

	// Response with a 200 OK status and the JSON forex record // as the body.
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Headers":     "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		},
	}, nil
}

// helper for sending responses relating to server errors.
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// helper for sending responses relating to client errors.
func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
func main() {
	lambda.Start(getForexAllItem)
}
