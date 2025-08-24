package notify

import (
	"fmt"
)

const Version = "0.0.1"

func Notify(message string) error {
	fmt.Printf("notification: %s\n", message)
}
