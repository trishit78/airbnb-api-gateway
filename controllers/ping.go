package controllers

import "net/http"

func PingHandler(w http.ResponseWriter, r *http.Request) {      // request is the same, response obj is diff every time

	w.Write([]byte ("pong"))
}