package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/info", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Host: ", req.Host)
		fmt.Fprintln(w, "URI: ", req.RequestURI)
		fmt.Fprintln(w, "Method: ", req.Method)
		fmt.Fprintln(w, "RemoteAddr: ", req.RemoteAddr)
	})

	http.ListenAndServe(":8080", nil)

}
