package main

import (
	"aws/instance"
	"aws/manager"
	"fmt"
	"os"
)

/**
go run aws.go terminateall & go run aws.go stopall
*/
func main() {
	reservations := instance.GetFilteredinstances()

	for _, each := range reservations.Reservations {
		for _, instance := range each.Instances {

			id := instance.InstanceId

			if os.Args[1] == "stopall" {
				manager.StopByIndex(id)
			}

			if os.Args[1] == "terminateall" {
				manager.TerminateByIndex(id)
			}

			fmt.Println(instance.InstanceId)
		}
	}

}
