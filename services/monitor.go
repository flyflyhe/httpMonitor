package services

import (
	"context"
	"encoding/json"
	"github.com/flyflyhe/httpMonitor"
	rpc2 "github.com/flyflyhe/httpMonitor/rpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rfyiamcool/go-timewheel"
	"github.com/rs/zerolog/log"
	"time"
	"unsafe"
)

type MonitorTask struct {
	*rpc2.UrlRequest
	IsAdd bool
}

var MonitorTaskChan chan *MonitorTask
var MonitorStart bool

type MonitorServer struct {
	rpc2.UnimplementedMonitorServerServer
	q        chan [2]string
	stopChan chan struct{}
	running  bool
	tw       *timewheel.TimeWheel
	tasks    map[string]*timewheel.Task
}

func (monitor *MonitorServer) start() error {
	var err error
	if !monitor.running {
		monitor.tw, err = timewheel.NewTimeWheel(1*time.Second, 360)
		if err != nil {
			return err
		}
		monitor.q = make(chan [2]string, 10)
		monitor.stopChan = make(chan struct{})
		monitor.running = true
		monitor.tasks = make(map[string]*timewheel.Task)
		MonitorStart = true
		MonitorTaskChan = make(chan *MonitorTask, 10)
	}

	return nil
}

func (monitor *MonitorServer) stop() error {
	monitor.tw.Stop()
	monitor.running = false
	MonitorStart = false
	//close(monitor.q) 防止数据未完全写入
	close(monitor.stopChan)
	close(MonitorTaskChan)
	return nil
}

func (monitor *MonitorServer) Start(req *rpc2.MonitorRequest, srv rpc2.MonitorServer_StartServer) error {
	if err := monitor.start(); err != nil {
		return err
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
				task := monitor.tw.AddCron(time.Duration(interval)*time.Millisecond, func() {
					addTaskFunc(url)
				})

				monitor.tasks[url] = task
			}

			monitor.tw.Start()
		}()
	}
	for {
		select {
		case mData := <-monitor.q:
			err := srv.Send(&rpc2.MonitorResponse{
				Result: mData[1],
				Url:    mData[0],
			})
			if err != nil {
				return err
			}
		case urlReq := <-MonitorTaskChan:
			if lastTask, ok := monitor.tasks[urlReq.Url]; ok {
				if err := monitor.tw.Remove(lastTask); err != nil {
					log.Error().Caller().Msg(err.Error())
				}
				log.Debug().Caller().Msg("remove task" + urlReq.String())
			}
			if urlReq.IsAdd {
				log.Debug().Caller().Msg("add task" + urlReq.String())
				task := monitor.tw.AddCron(time.Duration(urlReq.Interval)*time.Millisecond, func() {
					addTaskFunc(urlReq.Url)
				})
				monitor.tasks[urlReq.Url] = task
			}

		case <-monitor.stopChan:
			log.Debug().Caller().Uint32("address", uint32(uintptr(unsafe.Pointer(monitor)))).Msg("收到stop信号")
			return monitor.stop()
		default:
			time.Sleep(1 * time.Second) //防止频繁发送
			err := srv.Send(&rpc2.MonitorResponse{
				Result: "sleep",
			})
			if err != nil {
				return err
			}
		}
	}
}

func (monitor *MonitorServer) Stop(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	if monitor.running {
		monitor.stopChan <- struct{}{}
	}
	log.Debug().Caller().Uint32("address", uint32(uintptr(unsafe.Pointer(monitor)))).Msg("stop ")
	return &empty.Empty{}, nil
}
