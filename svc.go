package main

import (
	"context"
	"fmt"
	"math"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"

	spam_svc "github.com/StackAdapt/go-spam/external/services/grpc/definition"
	. "github.com/StackAdapt/go-spam/log"
)

var h prometheus.Histogram

func init() {
	h = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "sk_redis_delta_hist",
			Buckets: prometheus.ExponentialBuckets(1, 2, 20),
		})

}

type svc struct{}

func (svc) SetGet(ctx context.Context, req *spam_svc.Request) (*spam_svc.Response, error) {
	if req.Count <= 0 { // always increment by 1 at least
		req.Count += 1
	}

	var redis_count int
	err := pool.Do(ip_spam_script.Cmd(&redis_count, string(req.Key), `5`))
	if err != nil {
		Logger.Error("redis failed", zap.Error(err))
	}

	sk_estimate, err := Estimate(req.Key, req.Count)
	if err != nil {
		return nil, fmt.Errorf("error: %+v", err)
	}

	Logger.Info("estimate", zap.Int("redis", redis_count), zap.Uint64("sk", sk_estimate))

	h.Observe(math.Abs(float64(uint64(redis_count) - sk_estimate)))
	return &spam_svc.Response{Count: sk_estimate}, nil
}

// sk_redis_delta_hist_bucket{le="1"} 1.80148e+06
// sk_redis_delta_hist_bucket{le="2"} 1.8038e+06
// sk_redis_delta_hist_bucket{le="4"} 1.804036e+06
// sk_redis_delta_hist_bucket{le="8"} 1.804043e+06
// sk_redis_delta_hist_bucket{le="16"} 1.804043e+06
// sk_redis_delta_hist_bucket{le="32"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="64"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="128"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="256"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="512"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="1024"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="2048"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="4096"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="8192"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="16384"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="32768"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="65536"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="131072"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="262144"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="524288"} 1.804044e+06
// sk_redis_delta_hist_bucket{le="+Inf"} 2.097152e+06
// sk_redis_delta_hist_sum 5.406888261956859e+24
// sk_redis_delta_hist_count 2.097152e+06
