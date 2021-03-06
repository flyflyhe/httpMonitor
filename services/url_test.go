package services

import (
	"context"
	rpc2 "github.com/flyflyhe/httpMonitor/rpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestUrlService_GetAllDomainAndInterval(t *testing.T) {
	go Start(address)
	time.Sleep(1 * time.Second) //等待服务启动

	conn, err := GetRpcConn()
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	rpcClient := rpc2.NewUrlServiceClient(conn)
	if _, err = rpcClient.SetUrl(context.Background(), &rpc2.UrlRequest{Url: "a", Interval: 1}); err != nil {
		t.Error(err)
	}

	if _, err = rpcClient.SetUrl(context.Background(), &rpc2.UrlRequest{Url: "b", Interval: 2}); err != nil {
		t.Error(err)
	}

	if m, err := rpcClient.GetAllDomainAndInterval(context.Background(), &empty.Empty{}); err != nil {
		t.Error(err)
	} else {
		assert.Contains(t, m.UrlInterval, "a")
		assert.Contains(t, m.UrlInterval, "b")
		assert.Equal(t, m.UrlInterval["a"], int32(1))
		assert.Equal(t, m.UrlInterval["b"], int32(2))
	}
}

func GetRpcConn() (*grpc.ClientConn, error) {
	tlsCredentials, err := loadClientTLSCredentials()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tlsCredentials))
	return conn, err
}

func TestUrlList(t *testing.T) {

	go Start(address)

	time.Sleep(1 * time.Second) //等待服务启动

	conn, err := GetRpcConn()
	if err != nil {
		t.Error(err)
	}

	defer conn.Close()

	rpcClient := rpc2.NewUrlServiceClient(conn)
	url := "https://www.baidu.com"
	t.Log("获取rpcClient\n")
	if res, err := rpcClient.SetUrl(context.Background(), &rpc2.UrlRequest{Url: url, Interval: 1000}); err != nil {
		t.Log(err, "\n")
	} else {
		t.Log(res.Result, "\n")
		if res.Result != "ok" {
			t.Errorf("期望ok 获得%s", res.Result)
		}

		if res, err := rpcClient.GetAll(context.Background(), &empty.Empty{}); err != nil {
			t.Error(err.Error())
		} else {
			t.Log(res, "\n")

			urlExist := false
			for _, v := range res.Urls {
				if v == url {
					urlExist = true
					break
				}
			}

			if !urlExist {
				t.Errorf("期望%s 未取到", url)
			}

			if res, err := rpcClient.DeleteUrl(context.Background(), &rpc2.UrlRequest{Url: url}); err != nil {
				t.Error(err.Error())
			} else {
				if res.Result != "ok" {
					t.Errorf("期望ok 获取%s", res.Result)
				}
			}
		}
	}
}
