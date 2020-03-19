package api
import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const (
	AppID = "TODO: ADD APP ID"

	AppKey = "root"

	AppSecret = "root"
)

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (TokenClient, error) {
	client := warden.NewClient(cfg, opts...)
	client.Use(addToken())
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewTokenClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc api.proto


//添加GRPC签名认证
func addToken() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption, ) error {

		//将之前放入context中的metadata数据取出，如果没有则新建一个metadata
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		//key
		md.Append("key", AppKey)
		//secret
		md.Append("secret", AppSecret)

		//将metadata数据装入context中
		ctx = metadata.NewOutgoingContext(ctx, md)

		//使用带有追踪数据的context进行grpc调用．
		err := invoker(ctx, method, req, resp, cc, opts...)

		return err
	}
}