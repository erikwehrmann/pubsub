package pubsub

import (
	"context"
	"sync"
)

type Bus struct {
	mu          sync.RWMutex
	subscribers map[string][]chan any
}

func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[string][]chan any),
	}
}

// Subscribe registers a channel to a topic
func (b *Bus) Subscribe(topic string) <-chan any {
	ch := make(chan any, 10)

	b.mu.Lock()
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	b.mu.Unlock()

	return ch
}

// Publish sends an event to all subscribers
func (b *Bus) Publish(ctx context.Context, topic string, event any) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subscribers[topic] {
		select {
		case ch <- event:
		case <-ctx.Done():
			return
		}
	}
}
