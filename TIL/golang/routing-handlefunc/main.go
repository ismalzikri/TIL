package main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("service running!"))
	}

	http.HandleFunc("/", handlerIndex)

	fmt.Println("server started at http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
