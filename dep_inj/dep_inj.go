package dep_inj

import (
	"log"
	"net/http"
	"io"
	"fmt"
)

func Greet(buffer io.Writer, greeting string) {
	fmt.Fprintf(buffer, "Hello, %s", greeting)
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func run() {
	log.Fatal(http.ListenAndServe(":5003", http.HandlerFunc(GreetingHandler)))
}