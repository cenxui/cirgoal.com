
/*
   Copyright 2010-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at
    http://aws.amazon.com/apache2.0/
   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Receive message from Queue with long polling enabled.
//
// Usage:
//    go run sqs_longpolling_receive_message.go -n queue_name -t timeout
func main() {
	var timeout int64
	flag.Int64Var(&timeout, "t", 20, "(Optional) Timeout in seconds for long polling")
	flag.Parse()

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	// Create a SQS service client.
	svc := sqs.New(sess)

	qURL := "https://sqs.ap-northeast-1.amazonaws.com/208853302302/sqs"


	// Receive a message from the SQS queue with long polling enabled.
	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &qURL,
		AttributeNames: aws.StringSlice([]string{
			"SentTimestamp",
		}),
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: aws.Int64(timeout),
	})
	if err != nil {
		exitErrorf("Unable to receive message from queue %q, %v.", err)
	}

	fmt.Printf("Received %d messages.\n", len(result.Messages))
	if len(result.Messages) > 0 {
		fmt.Println(result.Messages)
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}