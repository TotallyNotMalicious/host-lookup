package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var host string

	fmt.Println("Host: ")
	fmt.Scanln(&host) // get the goodies

	url := "http://ip-api.com/line/" + host
	transport := &http.Transport{
		ResponseHeaderTimeout: 5 * time.Second,
		DisableCompression:    true,
	}
	client := &http.Client{Transport: transport, Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	respstring := string(data)
	fmt.Println("\n" + respstring + "")
}
