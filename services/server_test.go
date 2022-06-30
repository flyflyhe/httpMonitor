package services

import (
	"context"
	"github.com/flyflyhe/httpMonitor/rpc"
	"io"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:50051"
)

func TestSimple(t *testing.T) {

	go Start(address)

	time.Sleep(2 * time.Second) //等待服务启动
	conn, err := GetRpcConn()
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建发送结构体
	req := rpc.SimpleRequest{
		Data: "stream server grpc ",
	}
	grpcClient := rpc.NewStreamServerClient(conn)
	// 调用我们的服务(ListValue方法)
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		t.Fatalf("Call ListStr err: %v", err)
	}
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		t.Log(res.StreamValue, "\n")
	}

	time.Sleep(5 * time.Second)

}

func TestMonitor(t *testing.T) {

	go Start(address)

	time.Sleep(2 * time.Second) //等待服务启动
	conn, err := GetRpcConn()
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	rpcClient := rpc.NewMonitorServerClient(conn)
	urlRpcClient := rpc.NewUrlServiceClient(conn)
	monitorStream, err := rpcClient.Start(context.Background(), &rpc.MonitorRequest{Operate: "start"})
	if err != nil {
		t.Log(err)
	} else {
		i := 0
		for {
			if i == 5 {
				_, _ = urlRpcClient.SetUrl(context.Background(), &rpc.UrlRequest{Url: "https://www.zhihu.com", Interval: 1000})
				_, _ = urlRpcClient.SetUrl(context.Background(), &rpc.UrlRequest{Url: "https://www.baidu.com", Interval: 1000})
			}

			if i == 100 {
				_, _ = urlRpcClient.DeleteUrl(context.Background(), &rpc.UrlRequest{Url: "https://www.zhihu.com"})
				_, _ = urlRpcClient.DeleteUrl(context.Background(), &rpc.UrlRequest{Url: "https://www.baidu.com"})
			}

			res, err := monitorStream.Recv()

			if err != nil {
				t.Log(err.Error())
				break
			}

			if err == io.EOF {
				t.Log("break")
				break
			}

			// 打印返回值
			t.Log(res.GetUrl())
			t.Log(res.GetResult())
			i++
		}
	}
}

func TestUrlClient(t *testing.T) {

	go Start(address)

	time.Sleep(1 * time.Second) //等待服务启动
	conn, err := GetRpcConn()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	rpcClient := rpc.NewUrlServiceClient(conn)
	log.Println("获取rpcClient")
	if res, err := rpcClient.SetUrl(ctx, &rpc.UrlRequest{Url: "https://www.baidu.com", Interval: 1000}); err != nil {
		log.Println(err)
	} else {
		log.Println(res.GetResult())
	}

	time.Sleep(2 * time.Second)
}
