package main

import (
	"bytes"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"time"
	//"math/rand"
)

func main() {
	fmt.Println("************** Starting *****************")

	// url := "http://AZSDL-WL5CG9036P64:8090/v1/teachers"
	//url := "http://ptsv2.com/t/paulelong/post"


	// buf := make([]byte, 262143, 2000000)
	buf := make([]byte, 8000000, 8000000)

	for i := range buf {
		buf[i] = byte(i) //byte(rand.Intn(244))
	}

	for j := 0; j <= 24; j++ {
		fmt.Println("Sending block ", j)
		SendReq(buf)
		fmt.Println("Sent block ", j)
	}
}

func SendReq(data []byte) {
	// url := "http://b4f19a2a78edbf40d68269381026cf4d.m.pipedream.net"
	url := "http://AZSDL-WL5CG9036P64:8090"
	
	reader := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Host", "ptsv2.com")

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 100}

	// Validate cookie and headers are attached
	fmt.Println(req.Cookies())
	fmt.Println(req.Header)

	fmt.Println("************** Sending Data *****************")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal("Error reading body. ", err)
	// }

	// fmt.Println("************** Body is *****************")
	// fmt.Printf("%s\n", body)
}
