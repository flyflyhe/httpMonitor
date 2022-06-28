package services

import (
	rpc2 "github.com/flyflyhe/httpMonitor/rpc"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

// SimpleService 定义我们的服务

type StreamService struct {
	rpc2.UnimplementedStreamServerServer
}

// ListValue 实现ListValue方法
func (s *StreamService) ListValue(req *rpc2.SimpleRequest, srv rpc2.StreamServer_ListValueServer) error {
	for n := 0; n < 500; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		if err := srv.Send(&rpc2.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		}); err != nil {
			return err
		}

		log.Debug().Caller().Int32("call", int32(n)).Send()
		time.Sleep(1 * time.Second)
	}
	return nil
}
