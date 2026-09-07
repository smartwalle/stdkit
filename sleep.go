package stdkit

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, duration time.Duration) error {
	var timer = time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
