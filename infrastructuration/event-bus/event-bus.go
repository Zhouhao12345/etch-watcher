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
	Serialize() (string, error)
}

type Consumer interface {
	AcceptEvent(e CustomEvent) error
	GetPrivilege() string
}

type EventBus struct {
	high chan CustomEvent
	mid chan CustomEvent
	low chan CustomEvent
	highConsumers []Consumer
	midConsumers []Consumer
	lowConsumers []Consumer
}

func NewEventBus() *EventBus {
	return &EventBus{
		high:make(chan CustomEvent, 1000),
		mid:make(chan CustomEvent, 1000),
		low:make(chan CustomEvent, 1000),
		highConsumers: make([]Consumer, 0),
		midConsumers: make([]Consumer, 0),
		lowConsumers: make([]Consumer, 0),
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

func (eb *EventBus) Register(c Consumer, privilege string) error {
	switch privilege {
	case HIGH:
		eb.highConsumers = append(eb.highConsumers, c)
	case MID:
		eb.midConsumers = append(eb.midConsumers, c)
	case LOW:
		eb.lowConsumers = append(eb.lowConsumers, c)
	default:
		panic(privilege)
	}
	return nil
}

func (eb *EventBus) Push2Consumer(e CustomEvent, privilege string) {
	switch privilege{
	case HIGH:
		for _, c:= range eb.highConsumers {
			if err:= c.AcceptEvent(e); err!= nil {
				panic(err)
			}
		}
	case MID:
		for _, c:=range eb.midConsumers {
			if err:= c.AcceptEvent(e); err != nil {
				panic(err)
			}
		}
	case LOW:
		for _,c:=range eb.lowConsumers {
			if err := c.AcceptEvent(e); err != nil {
				panic(err)
			}
		}
	default:
		panic(privilege)
	}
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

