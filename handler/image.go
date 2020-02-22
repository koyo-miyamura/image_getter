package handler

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/koyo-miyamura/image_server/client"
)

// ImageHandler is XXX
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rawURL := r.FormValue("url")

	client, err := client.NewClient(rawURL)
	if err != nil {
		log.Println("error newClient", err.Error())
	}

	res, err := client.Do()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// w.Header().Set("Access-Control-Allow-Origin", "*")

	e := base64.NewEncoder(base64.StdEncoding, w)
	e.Write(res)
	e.Close()

	return
}
