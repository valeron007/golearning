package main

import "net/http"

type Handler interface {
	ServerHTTP(w http.ResponseWriter, r *http.Request)
}
