//go:build windows

package notify

import (
	"fmt"
	"github.com/go-toast/toast"
)

func Send(title, message string) error {
	fmt.Printf("notification: %s\n", message)
	n := toast.Notification{
		Title:   title,
		Message: message,
	}
	err := n.Push()
	if err != nil {
		return Fatal(err)
	}
	return nil
}
