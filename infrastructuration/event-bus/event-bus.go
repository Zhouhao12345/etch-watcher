package event_bus

type CustomEvent interface {
	GetToken() string
	GetContent() string
	GetType() string
	String() string
}

type EventBus struct {
	bus chan CustomEvent
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus:make(chan CustomEvent),
	}
}

func (eb *EventBus) Publish(e CustomEvent) error {
	eb.bus <- e
	return nil
}

func (eb *EventBus) Accept() chan CustomEvent{
	return eb.bus
}

