package services

import (
	"fmt"
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
				fmt.Print(event)
			default:
				select {
				case event =<-Bus.Accept(event_bus.HIGH):
					fmt.Print(event)
				case event =<-Bus.Accept(event_bus.MID):
					fmt.Print(event)
				default:
					select {
					case event =<-Bus.Accept(event_bus.HIGH):
						fmt.Print(event)
					case event =<-Bus.Accept(event_bus.MID):
						fmt.Print(event)
					case event =<-Bus.Accept(event_bus.LOW):
						fmt.Print(event)
					}
				}
			}
		}
	}()
}
