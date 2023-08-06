package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	url := "https://www.qq.com/"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	reqest.Header.Set("Retry-After", "120")
	reqest.Header.Set("Cookie", "120")
	reqest.Header.Set("User-Agent", "120")
	fmt.Printf("header: %+v\n", reqest.Header)

	rsp, _ := client.Do(reqest)
	rsp.Body.Close()

	fmt.Printf("rsp: %+v\n", rsp)

	_, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Printf("ReadAll err: %s\n", err)
		return
	}
}
