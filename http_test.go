package httpMonitor

import "testing"

func Test_send(t *testing.T) {
	err := send("https://www.baidu.com", "")

	if err != nil {
		t.Error(err)
	}
}
