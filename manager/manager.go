package manager

import (
	"aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ec2svc = session.CreateEc2Session()

func StopByIndex(id *string) {
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(*id),
		},
		DryRun: aws.Bool(true),
	}

	result, err := ec2svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = ec2svc.StopInstances(input)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("Success", result.StoppingInstances)
		}
	} else {
		fmt.Println("Error", err)
	}
}

func TerminateByIndex(id *string) {
	input := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{
			aws.String(*id),
		},
		DryRun: aws.Bool(true),
	}

	result, err := ec2svc.TerminateInstances(input)
	awsErr, ok := err.(awserr.Error)
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = ec2svc.TerminateInstances(input)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("Success", result.TerminatingInstances)
		}
	} else {
		fmt.Println("Error", err)
	}
}
