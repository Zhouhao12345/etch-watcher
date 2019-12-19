package event_bus

import (
	"fmt"
)

const (
	HIGH = "high"
	MID = "mid"
	LOW = "low"
)

type CustomEvent interface {
	GetToken() string
	GetContent() string
	GetType() string
	String() string
}

type EventBus struct {
	high chan CustomEvent
	mid chan CustomEvent
	low chan CustomEvent
}

func NewEventBus() *EventBus {
	return &EventBus{
		high:make(chan CustomEvent, 1000),
		mid:make(chan CustomEvent, 1000),
		low:make(chan CustomEvent, 1000),
	}
}

func (eb *EventBus) Publish(e CustomEvent, privilege string) error {
	defer func() {
		if r:= recover(); r!=nil{
			fmt.Printf("Recoverd for Privilege: %s", privilege)
		}
	}()
	switch privilege {
	case HIGH:
		eb.high <-e
	case MID:
		eb.mid <- e
	case LOW:
		eb.low <- e
	default:
		panic(privilege)
	}
	return nil
}

func (eb *EventBus) Accept(privilege string) chan CustomEvent{
	defer func() {
		if r:= recover(); r!=nil{
			fmt.Printf("Recoverd for Privilege: %s", privilege)
		}
	}()
	switch privilege {
	case HIGH:
		return eb.high
	case MID:
		return eb.mid
	case LOW:
		return eb.low
	default:
		panic(privilege)
	}
}

