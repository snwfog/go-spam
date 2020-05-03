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
