package gapi

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	starTime := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(starTime)

	statusCOde := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCOde = st.Code()
	}

	logger := log.Info()
	if err != nil {
		logger = log.Error().Err(err)
	}

	logger.Str("protocol", "grpc").
		Str("method", info.FullMethod).
		Int("status_code", int(statusCOde)).
		Str("status_text", statusCOde.String()).
		Dur("duration", duration).
		Msg("received a gRPC request")

	return result, err
}
