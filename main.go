package main

import "fmt"

func SystemRunningMessage() string {
	const message string = "System is running"
	return message
}

func main() {
	message := SystemRunningMessage()
	fmt.Println(message)

	return
}
