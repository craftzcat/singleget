package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {

}

func get() {
	// Create a query string.
	values := url.Values{
		"query": {"hello world"},
	}

	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
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

// this method can't get the body.
func head() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
}

// Send the POST method in x-www-form-urlencoded format.
func postFrom() {
	// Convert a string to the io.Reader interface.
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

// Send a string using the POST method.
func post() {
	reader := strings.NewReader("テキスト")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func postFile() {
	// Read a file in io.Reader format.
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func postMultiport() {
	// Declare a buffer to store a string of bytes after assembling the multipart part.
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo,jpg")
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}

	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
