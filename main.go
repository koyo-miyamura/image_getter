package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/koyo-miyamura/image_getter/handler"
)

const defaultPort = 12345

func main() {
	port, err := getPort()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Server started! Port:%d", port)

	server := handler.NewServer()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}

func getPort() (int, error) {
	var (
		portStr = os.Getenv("PORT")
	)
	if portStr == "" {
		return defaultPort, nil
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Println(err)
		return 0, nil
	}
	return port, nil
}
