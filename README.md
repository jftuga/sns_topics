# sns_topics
Command line tool for working with AWS SNS topics

## Current Status
* Only one operation is currently implemented:
* * List all AWS SNS topics by simultaneously querying all regions at once

## Example
```
$ sns_topics

arn:aws:sns:us-east-1:123456789012:alpha
arn:aws:sns:us-east-1:123456789012:beta
arn:aws:sns:us-east-1:123456789012:submit_new
arn:aws:sns:us-east-2:123456789012:submit_old
arn:aws:sns:us-west-2:123456789012:test_1
arn:aws:sns:ap-south-1:123456789012:test_2
```
