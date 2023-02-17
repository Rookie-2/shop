package register

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net/http"
	"shop/lib/logger"
)

type GrpcGatewayConfig struct {
	Name         string
	Addr         string
	RegisterFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
}

func RunGrpcGateway(ctx context.Context, gateConfigs []*GrpcGatewayConfig) error {
	// 1. 创建路由
	mux := runtime.NewServeMux()

	for _, cfg := range gateConfigs {
		// 2. 注册请求服务端
		if err := cfg.RegisterFunc(ctx, mux, cfg.Addr, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
			logger.Errorw("register grpc gateway failed", "ServerName", cfg.Name, "Error", err)
			return err
		}
		// 3. 启动并监听http请求
		if err := http.ListenAndServe(cfg.Addr, mux); err != nil {
			logger.Errorw("grpc gateway listenAndServer failed", "ServerName", cfg.Name, "Error", err)
			return err
		}
	}
	return nil
}
