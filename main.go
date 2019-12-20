package main

import (
	"fmt"
	"math"
	"time"
	_ "zhouhao.com/elevator/application/services/config"
	_ "zhouhao.com/elevator/application/services/logger"
	_ "zhouhao.com/elevator/application/services/event-bus"
	_ "zhouhao.com/elevator/application/services/redis"
)

func main() {
	var (
		start time.Time
	)
	start = time.Now()
	for {
			cur := time.Now()
			dur := cur.Sub(start)
			durSecs := dur.Seconds()
			durSecsRound := math.Floor(durSecs)
			if durSecs-durSecsRound <= math.Pow(10, -6) {
				fmt.Printf("Wait For Seconds: %f\n", durSecsRound)
			}
		}
}
