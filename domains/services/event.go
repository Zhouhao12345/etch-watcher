package services

import (
	"context"
	event_bus "zhouhao.com/elevator/infrastructuration/event-bus"
)

type PersistService struct {
	con context.Context
}

func NewPersistService()  {
	
}

func (s *PersistService) SaveToCache(c event_bus.CustomEvent) error {
	
}
