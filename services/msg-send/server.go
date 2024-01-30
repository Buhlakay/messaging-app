package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Buhlakay/messaging-app/msg-send/database"
)

type sendServer struct {
	// logf controls where logs are sent.
	logf     func(f string, v ...interface{})
	serveMux http.ServeMux
}

type messageFormat struct {
	receiverId     int64
	senderUsername string
	body           string
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

	var m messageFormat

	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	fmt.Printf("rx id: %d, tx username: %s, body: %s", m.receiverId, m.senderUsername, m.body)
	database.WriteMessage(m.receiverId, m.senderUsername, m.body)

	w.WriteHeader(http.StatusAccepted)
}
