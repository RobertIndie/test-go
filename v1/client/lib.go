package client

import "fmt"

func Hello(name string, id int) string {
	return fmt.Sprintf("Hello, %s [%d]", name, id)
}
