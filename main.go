package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joshuarubin/go-sway"
)

const (
	prevIdentifier string = "_prev"
	sentinelValue  int64  = -1
)

var prevFocus int64 = sentinelValue

type handler struct {
	sway.EventHandler
	sway.Client
}

func (h handler) Window(ctx context.Context, e sway.WindowEvent) {
	switch e.Change {
	case sway.WindowFocus:
		if e.Container.Focused && e.Container.ID > 0 && *e.Container.Visible {
			if prevFocus > 0 && prevFocus != e.Container.ID {
				h.UpdateMark(ctx)
			}
			prevFocus = e.Container.ID
		}
	case sway.WindowClose:
		if prevFocus == e.Container.ID {
			h.UpdateMark(ctx)
			prevFocus = sentinelValue
		}
	}
}

func (h handler) UpdateMark(ctx context.Context) {
	cmd := fmt.Sprintf("[con_id=%d] mark --add --toggle %s", prevFocus, prevIdentifier)

	if _, err := h.RunCommand(ctx, cmd); err != nil {
		log.Printf("Sway command '%s' failed: %s\n", cmd, err)
	}
}

func NewHandler(ctx context.Context) (*handler, error) {
	client, err := sway.New(ctx)
	if err != nil {
		return nil, err
	}

	h := &handler{
		sway.NoOpEventHandler(),
		client,
	}

	return h, nil
}

func main() {
	ctx := context.Background()

	h, err := NewHandler(ctx)
	if err != nil {
		log.Fatalf("Cannot create handler: %s\n", err)
	}

	if err = sway.Subscribe(ctx, h, sway.EventTypeWindow); err != nil {
		log.Fatalf("Cannot subscribe to Sway socket: %#v\n", err)
	}
}
