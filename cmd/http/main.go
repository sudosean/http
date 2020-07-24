package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// using htt.Method Get
	fmt.Println("http practice in go")
	r1, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	body, err := ioutil.ReadAll(r1.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	r1.Body.Close()
  // -------------------------------------------
	// using htt.Method HEAD
	r2, err := http.Head("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	r2.Body.Close()
	fmt.Println(r2.Status)
	// -------------------------------------------
	// using NewRequest
	req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
	if err != nil {
		log.Panicln(err)
	}
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println("DELETE ------", resp.Status)
}
