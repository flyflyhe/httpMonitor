package main

import (
	"fmt"
	"github.com/rfyiamcool/go-timewheel"
	"httpMonitor"
	"time"
)

func main() {

	tw, err := timewheel.NewTimeWheel(1*time.Second, 360)
	if err != nil {
		panic(err)
	}

	tw.AddCron(1*time.Second, func() {
		fmt.Println("执行")
		result, err := httpMonitor.Monitor("https://www.baidu.com")
		if err != nil {
			fmt.Println(err)
			return
		}
		for k, v := range result {
			fmt.Println(k, "--", v)
		}
	})

	tw.Start()

	select {}
}
