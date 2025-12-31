package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erikwehrmann/pubsub/api"
	"github.com/erikwehrmann/pubsub/pubsub"
	"github.com/erikwehrmann/pubsub/subscribers"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	bus := pubsub.NewBus()

	// Subscribe workers
	loggerCh := bus.Subscribe("order.created")
	emailCh := bus.Subscribe("order.created")
	analyticsCh := bus.Subscribe("order.created")

	subscribers.StartLogger(ctx, loggerCh)
	subscribers.StartEmailSender(ctx, emailCh)
	subscribers.StartAnalytics(ctx, analyticsCh)

	server := &api.Server{Bus: bus}

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", server.CreateOrder)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("HTTP server running on :8080")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpServer.Shutdown(shutdownCtx)
}
