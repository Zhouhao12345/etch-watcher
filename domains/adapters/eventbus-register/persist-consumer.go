package eventbus_register

import (
	"zhouhao.com/elevator/application/services/logger"
	event_bus "zhouhao.com/elevator/infrastructuration/event-bus"
)

type PersistConsumer struct {
	privilege string
}

func NewPersistConsumer() *PersistConsumer {
	return &PersistConsumer{
		privilege:event_bus.HIGH,
	}
}

func (c *PersistConsumer) AcceptEvent(e event_bus.CustomEvent) error {
	if jsonStr, err:= e.Serialize(); err != nil {
		return err
	} else {
		logger.Logger.Println(jsonStr)
		return nil
	}
}
func (c *PersistConsumer) GetPrivilege() string{
	return c.privilege
}
