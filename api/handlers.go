package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/erikwehrmann/pubsub/events"
	"github.com/erikwehrmann/pubsub/pubsub"
	"github.com/google/uuid"
)

type Server struct {
	Bus *pubsub.Bus
}

func (s *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := events.OrderCreated{
		OrderID:   uuid.NewString(),
		Amount:    99.99,
		CreatedAt: time.Now(),
	}

	s.Bus.Publish(r.Context(), "order.created", order)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(order)
}
