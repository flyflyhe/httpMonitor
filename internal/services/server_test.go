package services

import (
	"context"
	"github.com/flyflyhe/httpMonitor/internal/rpc"
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

	// 创建发送结构体
	req := rpc.SimpleRequest{
		Data: "stream server grpc ",
	}
	grpcClient := rpc.NewStreamServerClient(conn)
	// 调用我们的服务(ListValue方法)
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.StreamValue)
	}

}

func TestMonitor(t *testing.T) {

	go Start(address)

	time.Sleep(2 * time.Second) //等待服务启动
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
	rpcClient := rpc.NewMonitorServerClient(conn)
	monitorStream, err := rpcClient.Monitor(ctx, &rpc.MonitorRequest{Operate: "start"})
	if err != nil {
		log.Println(err)
	} else {
		for {
			//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
			res, err := monitorStream.Recv()

			if err != nil {
				log.Println(err.Error())
			}
			// 打印返回值
			log.Println(res.GetUrl())
			log.Println(res.GetResult())
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
	rpcClient := rpc.NewUrlServiceClient(conn)
	log.Println("获取rpcClient")
	if res, err := rpcClient.SetUrl(ctx, &rpc.UrlRequest{Url: "https://www.baidu.com", Interval: 1000}); err != nil {
		log.Println(err)
	} else {
		log.Println(res.GetResult())
	}

	time.Sleep(2 * time.Second)
}
