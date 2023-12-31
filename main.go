package main

import (
	"fmt"
	"learnWithTests/countdown"
	"learnWithTests/greetings"
	"log"
	"net/http"
	"os"
	"time"
)

func SystemRunningMessage(port int) string {
	greetings.Greet(os.Stdout, "Great Guy")
	const message string = "System is running in the port "
	if port == 0 {
		port = 8000
	}

	return message + fmt.Sprint(port)
}

func main() {
	message := SystemRunningMessage(8000)
	fmt.Println(message)

	sleeper := &countdown.ConfigurableSleeper{Duration: 1 * time.Second, Slept: time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)

	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(greetings.GreeterHandler)))
}
