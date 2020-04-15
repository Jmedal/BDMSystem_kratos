package grpc

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	AppKey = "root"

	AppSecret = "root"
)

//检查GRPC请求签名认证
func checkToken() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		log.Info("method:", info.FullMethod)

		//通过metadata
		md, exist := metadata.FromIncomingContext(ctx)
		if !exist {
			return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
		}

		var (
			appKey string

			appSecret string
		)

		//取出数据
		if key, ok := md["key"]; ok {
			appKey = key[0]
		}
		if secret, ok := md["secret"]; ok {
			appSecret = secret[0]
		}

		//验证
		if appKey != AppKey || appSecret != AppSecret {
			return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
		}

		resp, err = handler(ctx, req)
		return
	}
}
