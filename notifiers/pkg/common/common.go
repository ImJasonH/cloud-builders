package common

import (
	"fmt"

	_ "cloud.google.com/go/pubsub"
)

func Common() {
	fmt.Println("common!")
}
