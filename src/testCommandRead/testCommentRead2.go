package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleWebhook2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	log.Println("server started")
	log.Println("sync check #1")
	log.Println("sync check #2")
	http.HandleFunc("/webhook", handleWebhook2)
	log.Fatal(http.ListenAndServe(":12345", nil))
}