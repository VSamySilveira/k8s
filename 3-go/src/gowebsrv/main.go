package main

import (
	"fmt"
	"net/http"
)

func greeting(msg string) string {
	return "<b>"+msg+"</b>"
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>"+greeting("Code.education Rocks!")+"</body></html>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port 8000.")
	http.ListenAndServe(":8000", nil)
}
