package client

import "fmt"

func Hello(name string, id int) (string, error) {
	return fmt.Sprintf("Hello, %s [%d]", name, id), nil
}
