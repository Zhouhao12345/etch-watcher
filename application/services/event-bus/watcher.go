package event_bus

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
	"zhouhao.com/elevator/application/services/config"
	"zhouhao.com/elevator/infrastructuration/etcd"
	"zhouhao.com/elevator/infrastructuration/event-bus"
)

func WatcherInit(bus *event_bus.EventBus) {
	defer func() {
		if r:= recover(); r!=nil{
			fmt.Printf("%s", r)
		}
	}()


	for _, key := range config.EtcdConfig.GetWatchKeys(){
		go func(key map[string]string) {
			var (
				client *clientv3.Client
				err error
				rep clientv3.WatchChan
				content clientv3.WatchResponse
				ev *clientv3.Event
			)

			defer func() {
				if r:= recover(); r!=nil{
					fmt.Printf("%s", r)
					fmt.Printf("Recoverd for Key: %s", key[etcd.Key])
				}
			}()
			if client, err = etcd.NewClient(clientv3.Config{
				Endpoints:   config.EtcdConfig.GetHosts(),
				DialTimeout: 5 * time.Second,
			}); err != nil {
				panic(err)
			}
			defer func() {
				if err:= client.Close(); err != nil {
					panic(err)
				}
			}()
			for {
				rep = client.Watch(context.Background(), key[etcd.Key])
				for content = range rep {
					for _, ev = range content.Events{
						_ = bus.Publish(etcd.NewInnerEvent(ev), key[etcd.Privilege])
					}
				}
			}
		}(key)
	}
}
