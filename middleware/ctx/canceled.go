package ctx

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

var RequestCanceledErr = errors.New(http.StatusConflict, "RequestCanceledErr", "request canceled")

// Canceled 用于处理请求取消或超时的情况
func Canceled(timeout time.Duration) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			st := time.Now()
			reply, err = handler(ctx, req)
			if ctxErr := ctx.Err(); ctxErr != nil {
				if errors.Is(ctxErr, context.Canceled) || (errors.Is(ctxErr, context.DeadlineExceeded) && time.Since(st) < timeout) {
					return nil, RequestCanceledErr
				}
				if errors.Is(ctxErr, context.DeadlineExceeded) {
					return nil, RequestCanceledErr
				}
			}
			if err != nil {
				grpcErr, ok := status.FromError(err)
				if ok {
					switch grpcErr.Code() {
					case codes.Canceled:
						return nil, RequestCanceledErr
					case codes.DeadlineExceeded:
						return nil, RequestCanceledErr
					}
				}
			}
			return reply, err
		}
	}
}
