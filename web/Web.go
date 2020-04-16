package web

import (
	"fmt"
	"net/http"
)

//mywebserver function

func index(writer http.ResponseWriter, Request *http.Request) {
	fmt.Fprintf(writer, "<H1>My Go Server <H1>")
}

@export
func WebServer() {

	PORT := "3000"
	http.HandleFunc("/", index)
	fmt.Println("Starting at Port :", PORT)
	http.ListenAndServe(":"+PORT, nil)

}
