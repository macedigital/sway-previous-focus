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
		isContainer := *e.Container.PID > 0 && *e.Container.Visible && e.Container.Type == "con"
		if isContainer && prevFocus != e.Container.ID {
			if prevFocus != sentinelValue {
				h.UpdateMark(ctx)
			}
			prevFocus = e.Container.ID
		}
	case sway.WindowClose:
		// presumably the container is closed already, remove mark ...
		if e.Container.ID == prevFocus {
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
