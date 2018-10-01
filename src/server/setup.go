package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RequestSummary struct {
	URL     string
	Method  string
	Headers http.Header
	Params  url.Values
	Auth    *url.Userinfo
	Body    string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rs := RequestSummary{
			URL:     r.URL.RequestURI(),
			Method:  r.Method,
			Headers: r.Header,
			Params:  r.URL.Query(),
			Auth:    r.URL.User,
			Body:    string(bytes),
		}

		resp, err := json.MarshalIndent(&rs, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resp)
		w.Write([]byte("\n"))
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("Exiting...")
}
