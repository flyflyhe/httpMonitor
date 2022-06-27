package services

import (
	"context"
	"encoding/json"
	"github.com/flyflyhe/httpMonitor"
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
)

type UrlService struct {
	rpc.UnimplementedUrlServiceServer
}

func (monitor *UrlService) SetUrl(c context.Context, request *rpc.UrlRequest) (*rpc.UrlResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			errMsg, _ := json.Marshal(err)
			log.Error().Caller().Msg(string(errMsg))
		}
	}()

	err := httpMonitor.SetUrl(request.GetUrl(), request.GetInterval())
	if MonitorStart && MonitorTaskChan != nil {
		MonitorTaskChan <- &MonitorTask{request, true}
	}
	return &rpc.UrlResponse{Result: "ok"}, err
}

func (monitor *UrlService) DeleteUrl(c context.Context, request *rpc.UrlRequest) (*rpc.UrlResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			errMsg, _ := json.Marshal(err)
			log.Error().Caller().Msg(string(errMsg))
		}
	}()
	err := httpMonitor.DeleteUrl(request.Url)
	if MonitorStart && MonitorTaskChan != nil {
		MonitorTaskChan <- &MonitorTask{request, false}
	}
	return &rpc.UrlResponse{Result: "ok"}, err
}

func (monitor *UrlService) SetProxy(c context.Context, request *rpc.ProxyRequest) (*rpc.ProxyResponse, error) {
	err := httpMonitor.SetUrlProxy(request.GetProxy())

	return &rpc.ProxyResponse{Result: "ok"}, err
}

func (monitor *UrlService) GetAll(c context.Context, _ *empty.Empty) (*rpc.UrlListResponse, error) {
	if urls, err := httpMonitor.GetAllUrls(); err != nil {
		log.Error().Caller().Msg(err.Error())
		return &rpc.UrlListResponse{}, err
	} else {
		urlList := make([]string, len(urls))
		i := 0
		for url, _ := range urls {
			urlList[i] = url
			i++
		}

		return &rpc.UrlListResponse{Urls: urlList}, nil
	}
}

func (monitor *UrlService) GetAllProxy(c context.Context, _ *empty.Empty) (*rpc.ProxyListResponse, error) {
	if proxyList, err := httpMonitor.GetAllProxyList(); err != nil {
		log.Error().Caller().Msg(err.Error())
		return &rpc.ProxyListResponse{}, err
	} else {
		return &rpc.ProxyListResponse{ProxyList: proxyList}, nil
	}
}

func (monitor *UrlService) DeleteProxy(c context.Context, request *rpc.ProxyRequest) (*rpc.ProxyResponse, error) {
	err := httpMonitor.DeleteUrl(request.Proxy)

	return &rpc.ProxyResponse{Result: "ok"}, err
}
