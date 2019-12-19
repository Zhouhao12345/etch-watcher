package main

import (
	"fmt"
	"math"
	"time"
	"zhouhao.com/elevator/application/services"
	_ "zhouhao.com/elevator/application/services"
)

func main() {
	var (
		start time.Time
	)
	start = time.Now()
	services.AppInit(services.Bus)
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
