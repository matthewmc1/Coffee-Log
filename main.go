package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	mux "github.com/gorilla/mux"
)

func main() {
	log.Println("Server is starting")

	l := log.New(os.Stdout, "coffee-api", 2)

	router := mux.NewRouter()
	router.NewRoute().Path("/coffee").Handler(NewCoffee(l)).Methods("GET", "POST")
	router.NewRoute().Path("/health").Handler(NewHealthCheck(l)).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         *&PORT,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	http.Handle("/", router)

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	log.Panicln("Receieved termination call", sig)

	tc, _ := context.WithTimeout(context.Background(), 3*time.Second)
	srv.Shutdown(tc)
}
