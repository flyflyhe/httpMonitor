package httpMonitor

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"github.com/flyflyhe/httpMonitor/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
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
	rpc.RegisterMonitorServiceServer(s, &services.UrlService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemClientCA, err := ioutil.ReadFile("cert/root.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/cert.pem", "cert/private.key")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}
