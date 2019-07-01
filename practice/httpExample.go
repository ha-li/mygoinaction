package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


// an example of how to process http pages in
// go, kinda crude example, but it works
func main() {

	uri := "https://google.com"
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("hi")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response")
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}