package services

import (
	"context"
	"encoding/json"
	"github.com/flyflyhe/httpMonitor"
	rpc2 "github.com/flyflyhe/httpMonitor/rpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
)

type UrlService struct {
	rpc2.UnimplementedUrlServiceServer
}

func (monitor *UrlService) SetUrl(c context.Context, request *rpc2.UrlRequest) (*rpc2.UrlResponse, error) {
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
	return &rpc2.UrlResponse{Result: "ok"}, err
}

func (monitor *UrlService) DeleteUrl(c context.Context, request *rpc2.UrlRequest) (*rpc2.UrlResponse, error) {
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
	return &rpc2.UrlResponse{Result: "ok"}, err
}

func (monitor *UrlService) SetProxy(c context.Context, request *rpc2.ProxyRequest) (*rpc2.ProxyResponse, error) {
	err := httpMonitor.SetUrlProxy(request.GetProxy())

	return &rpc2.ProxyResponse{Result: "ok"}, err
}

func (monitor *UrlService) GetAll(c context.Context, _ *empty.Empty) (*rpc2.UrlListResponse, error) {
	if urls, err := httpMonitor.GetAllUrls(); err != nil {
		log.Error().Caller().Msg(err.Error())
		return &rpc2.UrlListResponse{}, err
	} else {
		urlList := make([]string, len(urls))
		i := 0
		for url, _ := range urls {
			urlList[i] = url
			i++
		}

		return &rpc2.UrlListResponse{Urls: urlList}, nil
	}
}

func (monitor *UrlService) GetAllDomainAndInterval(c context.Context, _ *empty.Empty) (*rpc2.UrlIntervalResponse, error) {
	if urls, err := httpMonitor.GetAllUrls(); err != nil {
		log.Error().Caller().Msg(err.Error())
		return &rpc2.UrlIntervalResponse{}, err
	} else {

		return &rpc2.UrlIntervalResponse{UrlInterval: urls}, nil
	}
}

func (monitor *UrlService) GetAllProxy(c context.Context, _ *empty.Empty) (*rpc2.ProxyListResponse, error) {
	if proxyList, err := httpMonitor.GetAllProxyList(); err != nil {
		log.Error().Caller().Msg(err.Error())
		return &rpc2.ProxyListResponse{}, err
	} else {
		return &rpc2.ProxyListResponse{ProxyList: proxyList}, nil
	}
}

func (monitor *UrlService) DeleteProxy(c context.Context, request *rpc2.ProxyRequest) (*rpc2.ProxyResponse, error) {
	err := httpMonitor.DeleteProxy(request.Proxy)

	return &rpc2.ProxyResponse{Result: "ok"}, err
}
