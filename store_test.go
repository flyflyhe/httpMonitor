package httpMonitor

import (
	"fmt"
	"testing"
)

var monitoryUrl = "https://www.baidu.com"

func TestSetUrlProxy(t *testing.T) {
	proxy := "http://127.0.0.1:8000"
	err := SetUrlProxy(monitoryUrl, proxy)
	if err != nil {
		t.Error(err)
	}

	proxyArr, err := GetUrlProxyList(monitoryUrl)
	if err != nil {
		t.Error(err)
	}

	if proxyArr[0] != proxy {
		t.Error("未获取到正常值")
	}

	fmt.Println(proxyArr[0])
}

func TestMonitor(t *testing.T) {

}
