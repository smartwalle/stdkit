package stdkit_test

import (
	"context"
	"testing"
	"time"

	"github.com/smartwalle/stdkit"
)

func TestContext_Cancel(t *testing.T) {
	var ctx = stdkit.NewContext(context.Background())
	defer ctx.Cancel()

	go func() {
		time.Sleep(time.Second)
		t.Log("cancel...")
		ctx.Cancel()
	}()

	t.Log("waiting:", ctx.Cancelled())
	ctx.Wait()
	t.Log("done:", ctx.Cancelled())
}

func TestContext_Timeout(t *testing.T) {
	var ctx = stdkit.ContextWithTimeout(context.Background(), time.Second*5)
	defer ctx.Cancel()
	t.Log("waiting:", ctx.Cancelled())

	go func() {
		for {
			select {
			case <-ctx.Done():
			default:
				time.Sleep(time.Second)
				t.Log("process...")
			}
		}
	}()

	ctx.Wait()

	t.Log("done:", ctx.Cancelled(), ctx.Err(), ctx.Cause())
}

func TestContext_Deadline(t *testing.T) {
	var ctx = stdkit.ContextWithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer ctx.Cancel()
	t.Log("waiting:", ctx.Cancelled())

	go func() {
		for {
			select {
			case <-ctx.Done():
			default:
				time.Sleep(time.Second)
				t.Log("process...")
			}
		}
	}()

	ctx.Wait()

	t.Log("done:", ctx.Cancelled(), ctx.Err(), ctx.Cause())
}
