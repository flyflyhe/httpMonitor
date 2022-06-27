package services

import (
	"context"
	"encoding/json"
	"github.com/flyflyhe/httpMonitor"
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rfyiamcool/go-timewheel"
	"github.com/rs/zerolog/log"
	"time"
	"unsafe"
)

var MonitorTaskChan chan *rpc.UrlRequest
var MonitorStart bool

type MonitorServer struct {
	rpc.UnimplementedMonitorServerServer
	q     chan [2]string
	stop  chan struct{}
	start bool
	tw    *timewheel.TimeWheel
}

func (monitor *MonitorServer) Start(req *rpc.MonitorRequest, srv rpc.MonitorServer_StartServer) error {
	var err error
	if !monitor.start {
		monitor.tw, err = timewheel.NewTimeWheel(1*time.Second, 360)
		if err != nil {
			return err
		}
		monitor.q = make(chan [2]string, 10)
		monitor.stop = make(chan struct{})
		monitor.start = true
		MonitorStart = true
		MonitorTaskChan = make(chan *rpc.UrlRequest, 10)
	}

	addTaskFunc := func(url string) {
		if result, err := httpMonitor.Monitor(url); err != nil {
			log.Debug().Str("line 31", err.Error()).Send()
		} else {
			resultJson, _ := json.Marshal(result)
			monitor.q <- [2]string{url, string(resultJson)}
		}
	}

	if urls, err := httpMonitor.GetAllUrls(); err != nil {
		return err
	} else {
		go func() {
			for url, interval := range urls {
				//addCron 返回task 如果想加删除功能 需要保存url 与 task的关系
				monitor.tw.AddCron(time.Duration(interval)*time.Millisecond, func() {
					addTaskFunc(url)
				})
			}

			monitor.tw.Start()
		}()
	}
	for {
		select {
		case mData := <-monitor.q:
			err := srv.Send(&rpc.MonitorResponse{
				Result: mData[1],
				Url:    mData[0],
			})
			if err != nil {
				return err
			}
		case task := <-MonitorTaskChan:
			log.Debug().Caller().Msg("add task" + task.String())
			monitor.tw.AddCron(time.Duration(task.Interval)*time.Millisecond, func() {
				addTaskFunc(task.Url)
			})
		case <-monitor.stop:
			log.Debug().Caller().Uint32("address", uint32(uintptr(unsafe.Pointer(monitor)))).Msg("收到stop信号")
			monitor.tw.Stop()
			monitor.start = false
			MonitorStart = false
			close(monitor.q)
			close(monitor.stop)
			close(MonitorTaskChan)
			return nil
		default:
			time.Sleep(1 * time.Second) //防止频繁发送
			err := srv.Send(&rpc.MonitorResponse{
				Result: "sleep",
			})
			if err != nil {
				return err
			}
		}
	}
}

func (monitor *MonitorServer) Stop(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	if monitor.start {
		monitor.stop <- struct{}{}
	}
	log.Debug().Caller().Uint32("address", uint32(uintptr(unsafe.Pointer(monitor)))).Msg("stop ")
	return &empty.Empty{}, nil
}
