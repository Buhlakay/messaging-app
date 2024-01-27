package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Buhlakay/messaging-app/msg-send/database"
)

func run() error {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}
	log.Printf("listening on http://%v", l.Addr())

	ss := newSendServer()
	s := &http.Server{
		Handler:      ss,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	database.InitDb()

	errc := make(chan error, 1)
	go func() {
		errc <- s.Serve(l)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case err := <-errc:
		log.Printf("failed to serve: %v", err)
	case sig := <-sigs:
		log.Printf("terminating: %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return s.Shutdown(ctx)
}

func main() {
	log.SetFlags(0)

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
