package subscribers

import (
	"context"
	"log"
)

func StartLogger(ctx context.Context, ch <-chan any) {
	go func() {
		for {
			select {
			case event := <-ch:
				log.Printf("[LOGGER] Event received: %#v\n", event)

			case <-ctx.Done():
				log.Println("[LOGGER] shutting down")
				return
			}
		}
	}()
}
