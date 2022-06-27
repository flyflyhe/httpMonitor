package services

import (
	"context"
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestUrlList(t *testing.T) {

	go Start(address)

	time.Sleep(1 * time.Second) //等待服务启动
	tlsCredentials, err := loadClientTLSCredentials()
	if err != nil {
		t.Fatal("cannot load TLS credentials: ", err)
	}
	if err != nil {
		t.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	rpcClient := rpc.NewUrlServiceClient(conn)
	url := "https://www.baidu.com"
	t.Log("获取rpcClient\n")
	if res, err := rpcClient.SetUrl(context.Background(), &rpc.UrlRequest{Url: url, Interval: 1000}); err != nil {
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

			if res, err := rpcClient.DeleteUrl(context.Background(), &rpc.UrlRequest{Url: url}); err != nil {
				t.Error(err.Error())
			} else {
				if res.Result != "ok" {
					t.Errorf("期望ok 获取%s", res.Result)
				}
			}
		}
	}
}
