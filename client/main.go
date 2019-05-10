package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://localhost:8080/Hubert et moi")
	if err != nil {
		fmt.Printf("Connection error");
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read error");
	}
	err = nil
	fmt.Printf("%s\n", body)
}