package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.mycityinfo.co.za")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("%#v\n\n", resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	s := string(body)

	fmt.Println(s)
}
