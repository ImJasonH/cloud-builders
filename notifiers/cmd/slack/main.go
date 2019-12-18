package main

import (
	"fmt"

	"github.com/GoogleCloudPlatform/cloud-builders/notifiers/pkg/common"
)

func main() {
	common.Common()
	fmt.Println("I'm Slack")
}
