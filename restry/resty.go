package restry

import (
	conf "github.com/fzf-labs/fkratos-contrib/api/conf/v1"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"

	"github.com/dubonzi/otelresty"
	"github.com/go-kratos/kratos/v2/log"
)

// NewResty  http 客户端
func NewResty(cfg *conf.Bootstrap, logger log.Logger) *resty.Client {
	l := log.NewHelper(log.With(logger, "module", "resty"), log.WithMessageKey("message"))
	client := resty.New()
	// 设置请求头
	client.SetHeader("app-from", cfg.Name)
	// 设置超时时间
	client.SetTimeout(5 * time.Second)
	// 禁止重定向
	client.SetRedirectPolicy(resty.NoRedirectPolicy())
	// 设置最大重试次数
	client.SetRetryCount(2)
	// 设置重试等待时间  默认100 milliseconds
	client.SetRetryWaitTime(100 * time.Microsecond)
	// 设置重试最大等待时间 默认2s
	client.SetRetryMaxWaitTime(2 * time.Second)
	// 链路追踪
	opts := []otelresty.Option{otelresty.WithTracerName(cfg.Name)}
	otelresty.TraceClient(client, opts...)
	// 设置错误时触发
	client.OnError(func(r *resty.Request, err error) {
		if err != nil {
			l.WithContext(r.Context()).Warnw("message", "restyErr", "path", r.URL, "method", r.Method, "headers", r.Header, "query", r.QueryParam, "req_body", r.Body, "err", err.Error())
		}
	})
	// 设置结束后执行
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		// 非200状态码日志
		if r.RawResponse.StatusCode != http.StatusOK {
			l.WithContext(r.Request.Context()).Warnw("message", "restyNotOK", "path", r.Request.URL, "method", r.Request.Method, "headers", headerToMap(r.Request.Header), "query", r.Request.QueryParam, "req_body", r.Request.Body, "status", r.StatusCode(), "res_body", r.String(), "latency", r.Time().Milliseconds())
		}
		// 慢日志
		if r.Time() > time.Millisecond*500 {
			l.WithContext(r.Request.Context()).Warnw("message", "restySlow", "path", r.Request.URL, "method", r.Request.Method, "headers", headerToMap(r.Request.Header), "query", r.Request.QueryParam, "req_body", r.Request.Body, "status", r.StatusCode(), "res_body", r.String(), "latency", r.Time().Milliseconds())
		}
		return nil
	})
	return client
}

// headerToMap 转换header
func headerToMap(header http.Header) map[string]string {
	result := make(map[string]string)
	for key, values := range header {
		if len(values) > 0 {
			result[key] = values[0]
		}
	}
	return result
}
