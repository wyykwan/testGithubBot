package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleWebhook3(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	log.Println("server started")
	log.Println("feature merge check #1")
	log.Println("feature merge check #2")
	log.Println("feature merge check #3??")
	http.HandleFunc("/webhook", handleWebhook3)
	log.Fatal(http.ListenAndServe(":12345", nil))
}