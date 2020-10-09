package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/blugelabs/bluge"
	"github.com/blugelabs/bluge/index"
	"github.com/blugelabs/bluge_directory_elf"
)

var reader *bluge.Reader

func internalError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       err.Error(),
	}, err
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var searchRequest SearchRequest
	err := json.Unmarshal([]byte(request.Body), &searchRequest)
	if err != nil {
		return internalError(fmt.Errorf("error unmarshaling search request: %v", err))
	}

	blugeRequest, err := searchRequest.BlugeRequest()
	if err != nil {
		return internalError(fmt.Errorf("error building bluge request: %v", err))
	}

	dmi, err := reader.Search(context.Background(), blugeRequest)
	if err != nil {
		return internalError(fmt.Errorf("error executing bluge search: %v", err))
	}

	searchResponse, err := NewSearchResponse(searchRequest.Query, dmi)
	if err != nil {
		return internalError(fmt.Errorf("error processing bluge response: %v", err))
	}
	searchResponse.AddPaging(dmi.Aggregations(), searchRequest.Page)

	responseBytes, err := json.Marshal(searchResponse)
	if err != nil {
		return internalError(fmt.Errorf("error marshaling search repsonse: %v", err))
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBytes),
	}, nil
}

func main() {
	cfg := bluge.DefaultConfigWithDirectory(func() index.Directory {
		return bluge_directory_elf.NewElfDirectory(os.Args[0], "index")
	})

	var err error
	reader, err = bluge.OpenReader(cfg)
	if err != nil {
		log.Fatalf("error opening index reader: %v", err)
	}

	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
