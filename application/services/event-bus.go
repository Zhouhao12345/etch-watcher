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
			case event = <-Bus.Accept():
				fmt.Print(event)
			}
		}
	}()
}
