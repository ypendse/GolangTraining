package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://google.com/"
	defer fmt.Printf("Completed getting: %s\n", url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("--------------------------------------")
		bs := make([]byte, 9999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
		fmt.Println("--------------------------------------")
		fmt.Printf("Return code: %d\n", resp.StatusCode)
		fmt.Println("--------------------------------------")
	}
}
