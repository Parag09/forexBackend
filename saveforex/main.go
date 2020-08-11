package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

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

// function to save data obtained from forex api to dynamo db database
func saveForexdata() ([]ForexData, error) {
	url := fmt.Sprintf("https://www.freeforexapi.com/api/live?pairs=EURGBP,USDJPY,USDCHF,GBPUSD,AUDUSD,NZDUSD,USDCAD,USDZAR")
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	var result ForexDataInput
	var output []ForexData

	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &result)

	for k, v := range result.Rates {
		val := v.(map[string]interface{})
		output = append(output, ForexData{
			Curr:      k,
			Rate:      fmt.Sprintf("%f", val["rate"]),
			Timestamp: fmt.Sprintf("%f", val["timestamp"]),
		})
	}
	// looping through the items got from forex api and calling put item
	for _, b := range output {
		err = putItem(b)
	}
	return output, err

}

func main() {

	lambda.Start(saveForexdata)
}
