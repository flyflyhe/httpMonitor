package services

import (
	"encoding/json"
	"github.com/flyflyhe/httpMonitor"
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"github.com/rfyiamcool/go-timewheel"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

var once sync.Once
var tw *timewheel.TimeWheel
var q chan [2]string

type MonitorServer struct {
	rpc.UnimplementedMonitorServerServer
}

func (monitor *MonitorServer) Monitor(req *rpc.MonitorRequest, srv rpc.MonitorServer_MonitorServer) error {
	q = make(chan [2]string, 10)
	if tw, err := GetTw(); err != nil {
		return err
	} else {
		if urls, err := httpMonitor.GetAllUrls(); err != nil {
			return err
		} else {
			go func() {
				for url, interval := range urls {
					tw.AddCron(time.Duration(int64(interval))*time.Millisecond, func() {
						if result, err := httpMonitor.Monitor(url); err != nil {
							log.Debug().Str("line 31", err.Error())
						} else {
							resultJson, _ := json.Marshal(result)
							q <- [2]string{url, string(resultJson)}
						}
					})
				}

				tw.Start()
			}()
		}
	}
	for {
		select {
		case mData := <-q:
			err := srv.Send(&rpc.MonitorResponse{
				Result: mData[1],
				Url:    mData[0],
			})
			if err != nil {
				return err
			}
		default:
			time.Sleep(1 * time.Second)
			err := srv.Send(&rpc.MonitorResponse{
				Result: "sleep",
			})
			if err != nil {
				return err
			}
		}
	}
}

func GetTw() (*timewheel.TimeWheel, error) {
	var err error
	once.Do(func() {
		tw, err = timewheel.NewTimeWheel(1*time.Second, 360)
	})
	return tw, err
}
