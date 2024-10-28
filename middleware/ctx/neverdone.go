package ctx

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"time"
)

// neverDoneCtx never done.
type neverDoneCtx struct {
	context.Context
}

// Done forbids the context done from parent context.
func (*neverDoneCtx) Done() <-chan struct{} {
	return nil
}

// Deadline forbids the context deadline from parent context.
func (*neverDoneCtx) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Err forbids the context done from parent context.
func (c *neverDoneCtx) Err() error {
	return nil
}

// NeverDone wraps and returns a new context object that will be never done,
// which forbids the context manually done, to make the context can be propagated
// to asynchronous goroutines.
//
// Note that, it does not affect the closing (canceling) of the parent context,
// as it is a wrapper for its parent, which only affects the next context handling.
func NeverDone() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = &neverDoneCtx{ctx}
			return handler(ctx, req)
		}
	}
}
