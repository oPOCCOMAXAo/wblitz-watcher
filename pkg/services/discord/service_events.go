package discord

import (
	"context"
)

type EventParams struct {
	Name    EventName
	Handler EventHandler
}

func (s *Service) RegisterEvent(params EventParams) {
	if params.Handler != nil {
		s.eventHandlers[params.Name] = params.Handler
	}
}

func (s *Service) processEvent(
	ctx context.Context,
	name EventName,
	data *Event,
) error {
	handler, ok := s.eventHandlers[name]
	if !ok {
		return nil
	}

	err := handler(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
