/*
sns_topics.go
-John Taylor
Sep-20-2021

A quick way to list all AWS SNS topics in all regions.

*/
package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"sort"
)

const version string = "1.1.1"

func main() {
	p := endpoints.AwsPartition()
	allRegions := getAllRegions(p.Regions())
	ch := make(chan []string)

	for _, region := range allRegions {
		go getTopicsInRegion(ch, region, false) // change to true to output error messages
	}

	var allTopics []string
	for i := 0; i < len(allRegions); i++ {
		regionTopics := <-ch
		if len(regionTopics) == 0 {
			continue
		}
		for _, topic := range regionTopics {
			allTopics = append(allTopics, topic)
		}
	}

	sort.Strings(allTopics)
	for _, arn := range allTopics {
		fmt.Println(arn)
	}
}

func getAllRegions(ep map[string]endpoints.Region) []string {
	var allRegions []string
	for name := range ep {
		allRegions = append(allRegions, name)
	}
	return allRegions
}

func getTopicsInRegion(ch chan []string, region string, showErrors bool) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	svc := sns.New(sess)
	result, err := svc.ListTopics(nil)
	if err != nil && showErrors {
		fmt.Println(region, err.Error())
	}

	var regionTopics []string
	for _, t := range result.Topics {
		regionTopics = append(regionTopics, *t.TopicArn)
	}
	ch <- regionTopics
}
