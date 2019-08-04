package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleWebhook4(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	log.Println("server started")
	log.Println("feature merge #1")
	log.Println("feature merge #2")
	http.HandleFunc("/webhook", handleWebhook4)
	log.Fatal(http.ListenAndServe(":12345", nil))
}