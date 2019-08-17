package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func wyyyyyyy(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	log.Println("server started")
	http.HandleFunc("/webhook", wyyyyyyy)
	log.Fatal(http.ListenAndServe(":12345", nil))
}