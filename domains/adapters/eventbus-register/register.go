package eventbus_register

import (
	event_bus2 "zhouhao.com/elevator/infrastructuration/event-bus"
)

func registerAllConsumer(bus  *event_bus2.EventBus,cons ...event_bus2.Consumer) error  {
	for _, con := range cons {
		if err := bus.Register(con ,con.GetPrivilege()); err != nil {
			return err
		}
	}
	return nil
}

func RegisterInit(bus *event_bus2.EventBus)  {
	registers := []event_bus2.Consumer{
		NewPersistConsumer(),
	}
	if err := registerAllConsumer(bus, registers...); err != nil {
		panic(err)
	}
}
