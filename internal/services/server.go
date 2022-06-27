package services

import (
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
)

func Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	c, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	//由于要发送较大的压缩包，默认为 4M。
	//如果需要向客户端发送大文件则增加一条grpc.MaxSendMsgSize()
	s := grpc.NewServer(
		grpc.Creds(c),
		grpc.MaxRecvMsgSize(math.MaxInt64))

	//注册服务
	rpc.RegisterUrlServiceServer(s, &UrlService{})
	rpc.RegisterMonitorServerServer(s, &MonitorServer{})
	rpc.RegisterStreamServerServer(s, &StreamService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
