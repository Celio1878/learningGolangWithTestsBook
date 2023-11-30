package main

import "fmt"

func SystemRunningMessage(port int) string {
	const message string = "System is running in the port "
	if port == 0 {
		port = 8000
	}

	return message + fmt.Sprint(port)
}

func main() {
	message := SystemRunningMessage(8000)
	fmt.Println(message)
}
