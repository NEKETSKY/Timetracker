package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type headers struct {
	Accept     []string `json:"Accept"`
	User_Agent []string `json:"User-Agent"`
}

type request struct {
	Host        string  `json:"host"`
	User_agent  string  `json:"user_agent"`
	Request_uri string  `json:"request_uri"`
	Headers     headers `json:"headers"`
}

func requesterHomework(w http.ResponseWriter, r *http.Request) {
	var rqst request
	rqst.Host = r.Host
	rqst.User_agent = r.UserAgent()
	rqst.Request_uri = r.RequestURI
	rqst.Headers.Accept = r.Header["Accept"]
	rqst.Headers.User_Agent = r.Header["User_Arent"]

	rqstJson, err := json.Marshal(rqst)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(rqstJson)
}

func main() {

	http.HandleFunc("/", requesterHomework)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}