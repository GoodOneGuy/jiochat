package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type RpcClient struct {
	C    interface{}
	Conn *grpc.ClientConn
	Ctx  context.Context
}

type NewFunc func(cc grpc.ClientConnInterface) interface{}

func NewRpcClient(addr string, fn NewFunc) *RpcClient {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return nil
	}

	// 新建一个客户端
	c := fn(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(
		ctx,
		"from", "test", "trace_id", "12345",
	)

	return &RpcClient{
		C:    c,
		Conn: conn,
		Ctx:  ctx,
	}
}

func (c *RpcClient) Close() {
	c.Conn.Close()
}
