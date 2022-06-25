package services

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:50051"
)

func TestClient(t *testing.T) {

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

func loadClientTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/root.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client_cert.pem", "cert/client_private.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		ServerName:   "test.com", //生成的证书通用名称 必须一致
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
