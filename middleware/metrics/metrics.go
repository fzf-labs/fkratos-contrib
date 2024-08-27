package metrics

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"go.opentelemetry.io/otel"
)

func Server(name string) middleware.Middleware {
	meter := otel.Meter(name)
	requests, err := metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
	if err != nil {
		panic(err)
	}
	seconds, err := metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
	if err != nil {
		panic(err)
	}
	return metrics.Server(
		metrics.WithRequests(requests),
		metrics.WithSeconds(seconds),
	)
}
