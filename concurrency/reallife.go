package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"http://www.rubyconf.com/",
	"http://golang.org",
	"http://matt.aimonetti.net",
}

type HttpResponse struct {
	url string
	response *http.Response
	err	error
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			ch <- &HttpResponse{url, resp, err}
			}(url)
		}
	}
	for {
		select {
		case r : <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return responses
}