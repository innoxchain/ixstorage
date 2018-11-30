package ixclient

import (
	"fmt"
)

// Greet is just a simple test function for other cosumers to
// test this package's functionality
func Greet(msg string) string {
	return fmt.Sprintf("Hello %s!\n", msg)
}