package api

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	"github.com/bilibili/kratos/pkg/net/rpc/warden/resolver"
	"google.golang.org/grpc/metadata"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const (
	AppID = "token.service"

	AppKey = "root"

	AppSecret = "root"
)

func init() {
	// NOTE: 注意这段代码，表示要使用discovery进行服务发现
	// NOTE: 还需注意的是，resolver.Register是全局生效的，所以建议该代码放在进程初始化的时候执行
	// NOTE: ！！！切记不要在一个进程内进行多个不同中间件的Register！！！
	// NOTE: 在启动应用时，可以通过flag(-discovery.nodes) 或者 环境配置(DISCOVERY_NODES)指定discovery节点
	resolver.Register(discovery.Builder())
}

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
//go:generate kratos tool protoc --grpc --bm api.proto

//生成 Swagger文档
//go:generate kratos tool protoc --swagger api.proto

//读取Swagger文档
//go:generate kratos tool swagger serve api.swagger.json --flavor=swagger
// --flavor=[redoc/swagger]

//添加GRPC签名认证
func addToken() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

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
