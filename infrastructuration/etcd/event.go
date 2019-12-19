package etcd

import "go.etcd.io/etcd/clientv3"

type InnerEvent struct {
	event *clientv3.Event
}

func NewInnerEvent(e *clientv3.Event) *InnerEvent {
	return &InnerEvent{event:e}
}

func (e *InnerEvent)GetType() string {
	return e.event.Type.String()
}

func (e *InnerEvent)GetToken() string {
	return string(e.event.Kv.Key)
}

func (e *InnerEvent)GetContent() string {
	return string(e.event.Kv.Value)
}

func (e *InnerEvent)String() string {
	return e.event.Kv.String()
}