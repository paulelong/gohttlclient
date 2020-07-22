package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	//"math/rand"
)

func main() {

	// url := "http://AZSDL-WL5CG9036P64:8090/v1/teachers"
	url := "http://ptsv2.com/t/paulelong/post"

	// buf := make([]byte, 262143, 2000000)
	buf := make([]byte, 200, 2000000)

	for i := range buf {
		buf[i] = byte(i) //byte(rand.Intn(244))
	}

	//d := bytes.NewBuffer(buf)
	reader := bytes.NewReader(buf)

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Host", "httpbin.org")

	// Create and Add cookie to request
	cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
	req.AddCookie(&cookie)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 100}

	// Validate cookie and headers are attached
	fmt.Println(req.Cookies())
	fmt.Println(req.Header)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("%s\n", body)
}
