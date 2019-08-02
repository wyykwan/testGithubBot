package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	log.Println("server started")
	log.Println("squash check #1")
	log.Println("squash check #2")
	log.Println("squash check #3")
	log.Println("test sync merge #1")
	log.Println("test sync merge #2")
	log.Println("test sync merge #3")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":12345", nil))
}