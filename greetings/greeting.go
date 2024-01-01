package greetings

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	res, err := fmt.Fprintf(writer, "Hey, %s!\n", name)

	if err != nil {
		fmt.Println(err, "Fprintf error")
	}

	fmt.Println(res, "- Number of bytes written")
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Friend!")
	fmt.Println(r.Body, "- Request Body")
}
