package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
	fmt.Fprint(writer, m)
}
func main() {
	msgHandler := msg("Hello from Csandra Service")
	fmt.Println("Server is listening...")
	http.ListenAndServe("Localhost:8181", msgHandler)
}
