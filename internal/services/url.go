package services

import (
	"context"
	"fmt"
	"github.com/flyflyhe/httpMonitor"
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"github.com/rfyiamcool/go-timewheel"
	"sync"
	"time"
)

var once sync.Once
var tw *timewheel.TimeWheel

type UrlService struct {
	rpc.UnimplementedUrlServiceServer
}

func (monitor *UrlService) SetUrl(c context.Context, request *rpc.UrlRequest) (*rpc.UrlResponse, error) {
	fmt.Println(request.GetUrl())
	err := httpMonitor.SetUrl(request.GetUrl(), request.GetInterval())
	return &rpc.UrlResponse{Result: "ok"}, err
}

func (monitor *UrlService) SetProxy(c context.Context, request *rpc.ProxyRequest) (*rpc.ProxyResponse, error) {
	fmt.Println(request.GetProxy())

	err := httpMonitor.SetUrlProxy(request.GetProxy())

	return &rpc.ProxyResponse{Result: "ok"}, err
}

func GetTw() (*timewheel.TimeWheel, error) {
	var err error
	once.Do(func() {
		tw, err = timewheel.NewTimeWheel(1*time.Second, 360)
	})
	return tw, err
}
