package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
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

	if err = WriteJSON(w, &ImageResponse{
		Base64: base64FromByte(res),
	}); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}

// ImageResponse は ImageHandler のレスポンスを表します
type ImageResponse struct {
	Base64 string `json:"base64"`
}

// WriteJSON はレスポンスをJSON形式で書き込みます
func WriteJSON(w http.ResponseWriter, res interface{}) error {
	result, err := json.Marshal(res)
	if err != nil {
		return err
	}

	_, err = w.Write(result)
	if err != nil {
		return err
	}
	return nil
}

func base64FromByte(b []byte) string {
	var (
		buf = &bytes.Buffer{}
		e   = base64.NewEncoder(base64.StdEncoding, buf)
	)
	e.Write(b)
	e.Close()

	return buf.String()
}
