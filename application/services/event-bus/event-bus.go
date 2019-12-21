package event_bus

import (
	"log"
	eventbus_register "zhouhao.com/elevator/domains/adapters/eventbus-register"
	"zhouhao.com/elevator/infrastructuration/event-bus"
)

var Bus *event_bus.EventBus

func init() {
	var (
		event event_bus.CustomEvent
		)
	Bus = event_bus.NewEventBus()
	go func() {
		for {
			select {
			case event = <-Bus.Accept(event_bus.HIGH):
				Bus.Push2Consumer(event, event_bus.HIGH)
			default:
				select {
				case event =<-Bus.Accept(event_bus.HIGH):
					Bus.Push2Consumer(event, event_bus.HIGH)
				case event =<-Bus.Accept(event_bus.MID):
					Bus.Push2Consumer(event, event_bus.MID)
				default:
					select {
					case event =<-Bus.Accept(event_bus.HIGH):
						Bus.Push2Consumer(event, event_bus.HIGH)
					case event =<-Bus.Accept(event_bus.MID):
						Bus.Push2Consumer(event, event_bus.MID)
					case event =<-Bus.Accept(event_bus.LOW):
						Bus.Push2Consumer(event, event_bus.LOW)
					}
				}
			}
		}
	}()
	eventbus_register.RegisterInit(Bus)
	WatcherInit(Bus)
	log.Print("EventBus Service Init Success")
}
