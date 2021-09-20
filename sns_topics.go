/*
sns_topics.go
-John Taylor
Sep-20-2021

A quick way to list all AWS SNS topics in all regions.

*/
package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"sync"

	"fmt"
)

const version string = "1.0.0"

var wg sync.WaitGroup

func main() {
	p := endpoints.AwsPartition()
	allRegions := getAllRegions(p.Regions())
	wg.Add(len(allRegions))

	for _, region := range allRegions {
		go getTopicsInRegion(region, false) // change to true to output error messages
	}
	wg.Wait()
}

func getAllRegions(ep map[string]endpoints.Region) []string {
	var allRegions []string
	for name := range ep {
		allRegions = append(allRegions, name)
	}
	return allRegions
}

func getTopicsInRegion(region string, showErrors bool) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	svc := sns.New(sess)
	result, err := svc.ListTopics(nil)
	if err != nil && showErrors {
		fmt.Println(region, err.Error())
	}

	for _, t := range result.Topics {
		fmt.Println(*t.TopicArn)
	}
	wg.Done()
}
