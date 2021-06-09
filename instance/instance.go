package instance

import (
	"aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
)

func GetFilteredinstances() *ec2.DescribeInstancesOutput {
	session := session.CreateEc2Session()

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String("runner-*")},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String("runner-*")},
			},
		},
	}

	result, err := session.DescribeInstances(input)

	if err != nil {
		fmt.Println("Error ", err)
	}

	if result.Reservations == nil {
		fmt.Println("No instances to stop left, exiting...")
		os.Exit(0)
	}

	return result
}
