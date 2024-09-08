package main

import (
	"fmt"
)

func main() {

	server := newAPIserver(":5000")
	server.Run()
	fmt.Println("hello")
}
