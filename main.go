package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(200)
	})
	fmt.Println("hello docker")

	http.ListenAndServe(":8181", nil)
}
