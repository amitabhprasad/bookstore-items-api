package controllers

import (
	"net/http"
)

var (
	PingController pingInterface = &pingController{}
)

const (
	pong = "pong"
)

type pingInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

func (p *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
