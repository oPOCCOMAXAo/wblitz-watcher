package app

import (
	"context"
	"time"
)

func (app *App) TaskClans(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	done := ctx.Done()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			app.ProcessClans(ctx)
		}
	}
}
