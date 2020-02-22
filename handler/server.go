package handler

import "net/http"

// NewServer is XXX
func NewServer() *http.ServeMux {
	server := http.NewServeMux()

	server.HandleFunc("/", ImageHandler)

	return server
}
