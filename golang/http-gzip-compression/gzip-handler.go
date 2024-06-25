// package main

// import (
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/NYTimes/gziphandler"
// )

// func main() {
// 	mux := new(http.ServeMux)

// 	mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
// 		f, err := os.Open("tes.jpeg")

// 		if f != nil {
// 			defer f.Close()
// 		}

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		_, err = io.Copy(w, f)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 	})

// 	server := new(http.Server)
// 	server.Addr = ":9000"
// 	server.Handler = gziphandler.GzipHandler(mux)

// 	server.ListenAndServe()
// }
