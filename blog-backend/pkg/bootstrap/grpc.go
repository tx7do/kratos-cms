package bootstrap

import (
	"context"
	"time"

	GRPC "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

const defaultTimeout = 5 * time.Second

func CreateGrpcClient(ctx context.Context, r registry.Discovery, serviceName string, timeoutDuration *durationpb.Duration) GRPC.ClientConnInterface {
	timeout := defaultTimeout
	if timeoutDuration != nil {
		timeout = timeoutDuration.AsDuration()
	}

	endpoint := "discovery:///" + serviceName

	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(r),
		grpc.WithTimeout(timeout),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Fatalf("dial grpc client [%s] failed: %s", serviceName, err.Error())
	}

	return conn
}
