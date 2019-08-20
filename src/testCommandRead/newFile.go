package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func hello (w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	log.Println("server starteddd")
	http.HandleFunc("/webhook", hello)
	log.Fatal(http.ListenAndServe(":12345", nil))
}