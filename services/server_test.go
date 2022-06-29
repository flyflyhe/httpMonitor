package services

import (
	"context"
	rpc2 "github.com/flyflyhe/httpMonitor/rpc"
	"google.golang.org/grpc"
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

	// 创建发送结构体
	req := rpc2.SimpleRequest{
		Data: "stream server grpc ",
	}
	grpcClient := rpc2.NewStreamServerClient(conn)
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

	rpcClient := rpc2.NewMonitorServerClient(conn)
	urlRpcClient := rpc2.NewUrlServiceClient(conn)
	monitorStream, err := rpcClient.Start(context.Background(), &rpc2.MonitorRequest{Operate: "start"})
	if err != nil {
		t.Log(err)
	} else {
		i := 0
		for {
			if i == 5 {
				urlRpcClient.SetUrl(context.Background(), &rpc2.UrlRequest{Url: "https://www.zhihu.com", Interval: 1000})
				urlRpcClient.SetUrl(context.Background(), &rpc2.UrlRequest{Url: "https://www.www.baidu.com", Interval: 1000})
			}

			if i == 100 {
				urlRpcClient.DeleteUrl(context.Background(), &rpc2.UrlRequest{Url: "https://www.zhihu.com"})
				urlRpcClient.DeleteUrl(context.Background(), &rpc2.UrlRequest{Url: "https://www.baidu.com"})
			}
			//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
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
	tlsCredentials, err := loadClientTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	rpcClient := rpc2.NewUrlServiceClient(conn)
	log.Println("获取rpcClient")
	if res, err := rpcClient.SetUrl(ctx, &rpc2.UrlRequest{Url: "https://www.baidu.com", Interval: 1000}); err != nil {
		log.Println(err)
	} else {
		log.Println(res.GetResult())
	}

	time.Sleep(2 * time.Second)
}
