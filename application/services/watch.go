package services

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
	c "zhouhao.com/elevator/config"
	"zhouhao.com/elevator/infrastructuration/etcd"
	"zhouhao.com/elevator/infrastructuration/event-bus"
)

func AppInit(bus *event_bus.EventBus) {
	defer func() {
		if r:= recover(); r!=nil{
			fmt.Printf("%s", r)
		}
	}()

	config := c.NewConfig()
	if err := c.LoadConfig(config); err != nil {
		panic(err)
	}
	for _, key := range config.Watch[c.Keys]{
		go func(key string) {
			defer func() {
				if r:= recover(); r!=nil{
					fmt.Printf("%s", r)
					fmt.Printf("Recoverd for Key: %s", key)
				}
			}()
			client, _ := clientv3.New(clientv3.Config{
				Endpoints:   config.Cluster[c.Hosts],
				DialTimeout: 5 * time.Second,
			})
			defer func() {
				if err:= client.Close(); err != nil {
					panic(err)
				}
			}()
			for {
				rep := client.Watch(context.Background(), key)
				for content := range rep {
					for _, ev := range content.Events{
						_ = bus.Publish(etcd.NewInnerEvent(ev))
					}
				}
			}
		}(key)
	}
}
