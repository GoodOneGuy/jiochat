package test

import (
	"fmt"
	"github.com/GoodOneGuy/jiochat/common/rpc"
	"github.com/GoodOneGuy/jiochat/proto/service"
	"google.golang.org/grpc"
	"testing"
)

func TestServer_SayHello(t *testing.T) {

	// 新建一个客户端

	c := rpc.NewRpcClient(":8000", func(cc grpc.ClientConnInterface) interface{} {
		return service.NewGreeterClient(cc)
	})
	defer c.Close()

	r, err := c.C.(service.GreeterClient).SayHello(c.Ctx, &service.HelloRequest{Name: "horika"})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}

	fmt.Printf("调用成功: %s", r.Message)
}
