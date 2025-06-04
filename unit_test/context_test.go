package main

import (
	"context"
	"testing"
	"time"
)

func work(ctx context.Context) error {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			// simulate work
		}
	}
	return nil
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	errChan := make(chan error)
	go func() {
		errChan <- work(ctx)
	}()

	select {
	case err := <-errChan:
		if err != context.DeadlineExceeded {
			t.Fatalf("expected context deadline exceeded, got %v", err)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatalf("timeout waiting for work to finish")
	}
}
