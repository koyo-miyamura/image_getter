package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const (
	// Timeout is XXX
	Timeout = 5 * time.Second
)

func main() {
	log.Println("Server started!")

	server := NewServer()

	log.Fatal(http.ListenAndServe(":12345", server))
}

// NewServer is XXX
func NewServer() *http.ServeMux {
	server := http.NewServeMux()

	server.HandleFunc("/", ImageHandler)

	return server
}

// ImageHandler is XXX
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rawURL := r.FormValue("url")

	client, err := NewClient(rawURL)
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

// Client is XXX
type Client struct {
	URL     *url.URL
	Timeout time.Duration
	client  *http.Client
}

// NewClient is XXX
func NewClient(urlStr string) (*Client, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("error Parse url %s", err.Error())
	}

	client := &http.Client{
		Timeout: Timeout,
	}

	return &Client{
		URL:     u,
		Timeout: Timeout,
		client:  client,
	}, nil
}

// Do is XXX
func (c *Client) Do() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.URL.String(), nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	var (
		b1 = &bytes.Buffer{}
		b2 = &bytes.Buffer{}
		mw = io.MultiWriter(b1, b2)
	)
	io.Copy(mw, res.Body)

	// 画像かどうか判定
	_, _, err = image.Decode(b1)
	if err != nil {
		return []byte{}, err
	}

	b, err := ioutil.ReadAll(b2)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
