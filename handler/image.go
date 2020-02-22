package handler

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/koyo-miyamura/image_getter/client"
)

// ImageHandler はリクエストされたURL先から画像を読み込んでbase64エンコーディングして返します
// URL先のコンテンツが画像でない場合はhttp.StatusBadRequestで返します
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rawURL := r.FormValue("url")

	c, err := client.NewClient(rawURL)
	if err != nil {
		log.Println("error newClient", err.Error())
	}

	res, err := c.Do()
	if err != nil {
		log.Println(err)
		if err == client.ErrInvalidImage {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	e := base64.NewEncoder(base64.StdEncoding, w)
	e.Write(res)
	e.Close()

	return
}
