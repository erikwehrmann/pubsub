package subscribers

import (
	"context"
	"log"
)

func StartAnalytics(ctx context.Context, ch <-chan any) {
	go func() {
		for {
			select {
			case event := <-ch:
				log.Println("[ANALYTICS] Event tracked:", event)
			case <-ctx.Done():
				log.Println("[ANALYTICS] Shutting down")
				return
			}
		}
	}()
}
