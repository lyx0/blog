package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/lyx0/blog/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()
	l := log.New(os.Stdout, "blog", log.LstdFlags)

	ah := handlers.NewArticle(l)

	r.HandleFunc("/", handlers.HomeHandler)
	r.Handle("/articles/{id:[0-9]+}", ah)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		logrus.Info("Listening on 8080")
		err := s.ListenAndServe()
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	logrus.Println("Received terminate", sig)

	tc, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		logrus.Fatal(err)
	}
	s.Shutdown(tc)

}
