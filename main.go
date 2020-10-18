package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}

	// Display as a string. (e.g., "200 OK")
	log.Println("Staus:", resp.Status)
	// Display as an integer. (e.g., 200)
	log.Println("StausCode:", resp.StatusCode)

	log.Println("Headers:", resp.Header)
	// Display a specific header.
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))

	defer resp.Body.Close()
	// Read the contents of io.Reader into a byte array.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
