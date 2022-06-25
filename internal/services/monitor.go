package services

import (
	"fmt"
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"github.com/rfyiamcool/go-timewheel"
	"sync"
	"time"
)

var once sync.Once
var tw *timewheel.TimeWheel

type MonitorServer struct {
	rpc.UnimplementedMonitorServerServer
}

func (monitor *MonitorServer) Monitor(req *rpc.MonitorRequest, srv rpc.MonitorServer_MonitorServer) error {
	fmt.Println(req.GetOperate())
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&rpc.MonitorResponse{
			Result: "oo",
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTw() (*timewheel.TimeWheel, error) {
	var err error
	once.Do(func() {
		tw, err = timewheel.NewTimeWheel(1*time.Second, 360)
	})
	return tw, err
}
