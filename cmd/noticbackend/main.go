package main

import (
	"context"
	"log"
	"noticbackend/app/server"
	"noticbackend/database"
	"os"
	"os/signal"
	"time"
)

func main() {

	srv := server.New()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(ch, os.Interrupt)

	// Block until we receive our signal.
	<-ch
	log.Println("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		if err := database.ShutDown(ctx); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()
	log.Println("server shut down")

	os.Exit(0)
}
