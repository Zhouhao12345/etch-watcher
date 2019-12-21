package main

import (
	"fmt"
	"time"
	_ "zhouhao.com/elevator/application/services/config"
	_ "zhouhao.com/elevator/application/services/event-bus"
	_ "zhouhao.com/elevator/application/services/logger"
	_ "zhouhao.com/elevator/application/services/redis"
)

func main() {
	var (
		start,oStart time.Time
		cur time.Time
	)
	oStart = time.Now()
	start = time.Now()
	for {
				cur = time.Now()
				durSecond := cur.Sub(start)
				if durSecond >= time.Second {
					start = cur
					oDurSecond := cur.Sub(oStart).Seconds()
					fmt.Printf("Wait For Seconds: %f\n", oDurSecond)
				}
		}
}
