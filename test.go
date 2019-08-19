package testmod

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func SayHi(name string) string {
	return fmt.Sprintf("Hi,%s", name)
}
