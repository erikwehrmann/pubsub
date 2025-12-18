package subscribers

import (
	"context"
	"log"
	"time"
)

func StartEmailSender(ctx context.Context, ch <-chan any) {
	go func() {
		for {
			select {
			case event := <-ch:
				log.Println("[EMAIL] Sending email for event", event)
				time.Sleep(500 * time.Millisecond)
				log.Println("[EMAIL] Email sent")

			case <-ctx.Done():
				log.Println("[EMAIL] shutting down")
				return
			}
		}
	}()
}
