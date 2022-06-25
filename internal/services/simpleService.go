package services

import (
	"github.com/flyflyhe/httpMonitor/internal/rpc"
	"strconv"
)

// SimpleService 定义我们的服务

type StreamService struct {
	rpc.UnimplementedStreamServerServer
}

// ListValue 实现ListValue方法
func (s *StreamService) ListValue(req *rpc.SimpleRequest, srv rpc.StreamServer_ListValueServer) error {
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&rpc.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
