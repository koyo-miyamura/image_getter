package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/koyo-miyamura/image_server/handler"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Please set port. ex) go run main.go 3000")
		return
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Server started! Port:%d", port)

	server := handler.NewServer()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
