package main

import (
	"io"
	"log"
	"net/http"
)

type sendServer struct {
	// logf controls where logs are sent.
	logf     func(f string, v ...interface{})
	serveMux http.ServeMux
}

func newSendServer() *sendServer {
	s := &sendServer{
		logf: log.Printf,
	}
	s.serveMux.Handle("/", http.FileServer(http.Dir(".")))
	s.serveMux.HandleFunc("/message", s.sendHandler)

	return s
}

func (s *sendServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.serveMux.ServeHTTP(w, r)
}

func (s *sendServer) sendHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body := http.MaxBytesReader(w, req.Body, 8192)
	_, err := io.ReadAll(body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusRequestEntityTooLarge), http.StatusRequestEntityTooLarge)
		return
	}

	// TODO: Write to DB
}
