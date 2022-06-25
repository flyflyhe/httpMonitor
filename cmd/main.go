package main

import (
	"flag"
	"fmt"
	"github.com/flyflyhe/httpMonitor"
	"github.com/rfyiamcool/go-timewheel"
	"log"
	"strings"
	"time"
)

/**
添加监控地址与代理地址
历史数据只要不删除monitor.db文件数据不会丢失
*/
func main() {
	interval := time.Duration(handleArgs())

	tw, err := timewheel.NewTimeWheel(1*time.Second, 360)
	if err != nil {
		panic(err)
	}

	urlArr, err := httpMonitor.GetByBucket(httpMonitor.BucketUrl)
	if err != nil {
		panic(err)
	}

	for _, url := range urlArr {
		tw.AddCron(interval*time.Second, func() {
			fmt.Println("执行监控:", url)
			result, err := httpMonitor.Monitor(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			for k, v := range result {
				fmt.Println(k, "--", v)
			}
		})
	}

	tw.Start()

	select {}
}

func handleArgs() int {
	var err error
	url := flag.String("url", "", "地址多个地址逗号分隔eg:https://www.baidu.com")
	proxy := flag.String("proxy", "", "代理地址,号分隔eg:socks5://127.0.0.1:8000")
	isDelete := flag.Bool("d", false, "是否删除 设置为是则删除所配置的url proxy")
	interval := flag.Int("i", 10, "时间间隔")
	flag.Parse()

	if *url != "" {
		urlList := strings.Split(*url, ",")
		for _, v := range urlList {
			v = strings.TrimSpace(v)
			if *isDelete {
				err = httpMonitor.Delete(httpMonitor.BucketUrl, v)
				log.Println("删除url:", v)
			} else {
				err = httpMonitor.SetUrl(v)
				log.Println("添加url:", v)
			}
			if err != nil {
				log.Println(err)
			}
		}
	}

	if *proxy != "" {
		proxyList := strings.Split(*proxy, ",")
		for _, v := range proxyList {
			v := strings.TrimSpace(v)
			if *isDelete {
				err = httpMonitor.Delete(httpMonitor.BucketProxy, v)
				log.Println("删除proxy:", v)
			} else {
				err = httpMonitor.SetUrlProxy(v)
				log.Println("添加proxy:", v)
			}
			if err != nil {
				log.Println(err)
			}
		}
	}

	return *interval
}
