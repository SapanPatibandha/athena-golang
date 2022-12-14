package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"

	// "github.com/aws/aws-sdk-go/aws/external"
	"github.com/aws/aws-sdk-go/service/athena"
)

const table = "textqldb.textqltable"
const outputBucket = "s3://business-time-nonprod"

func main() {
	cfg, err := external.LocalDefaultAWSConfig()

	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		return
	}

	cfg.Region = endpoints.UsEast2RegionID

	client := athena.New(cfg)

	query := "select * from " + table

	resultConf := &athena.ResultConfiguration{
		OutputLocation: aws.String(outputBucket),
	}

	params := &athena.StartQueryExecutionInput{
		QueryString:         aws.String(query),
		ResultConfiguration: resultConf,
	}

	// req := client.StartQueryExecutionRequest(params)

	req := client.GetQueryResultsPages(params)

	resp, err := req.Send(context.TODO())
	if err != nil {
		fmt.Printf("query error: %v\n", err)
		return
	}

	fmt.Println(resp)
}
