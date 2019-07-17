package modtest

import "fmt"

func Hello(name string) (string, error) {
	return fmt.Sprintf("v2: Hello %s!!!", name), nil
}
