package httpMonitor

import (
	"fmt"
	"testing"
)

var monitoryUrl = "https://www.baidu.com"

func TestSetUrlProxy(t *testing.T) {
	proxy := "http://127.0.0.1:8000"
	err := SetUrlProxy(proxy)
	if err != nil {
		t.Error(err)
	}

	result, err := GetByBucketAndKey(BucketProxy, proxy)
	if result != proxy {
		t.Errorf("期望%s获得%s", proxy, result)
	}

	if err = Delete(BucketProxy, proxy); err != nil {
		t.Errorf(err.Error())
	}
}

func TestMonitor(t *testing.T) {
	result, err := Monitor(monitoryUrl)
	if err != nil {
		t.Errorf(err.Error())
	}
	for k, v := range result {
		fmt.Println(k, "--", v)
	}
}
