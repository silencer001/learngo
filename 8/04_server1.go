package main

import (
	"fmt"
	"net/http"
)

func HandleConn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request url", r.URL)
	fmt.Println("request method:", r.Method)
	fmt.Println("request body:", r.Body)
	fmt.Println("request header:", r.Header)
	w.Write([]byte("hello go"))
}
func main() {
	http.HandleFunc("/go", HandleConn)

	http.ListenAndServe(":8000", nil)
}