package main

import (
	"net"
	"net/http"

	"github.com/mediocregopher/radix/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	spam_svc "github.com/StackAdapt/go-spam/external/services/grpc/definition"
	. "github.com/StackAdapt/go-spam/log"
)

var pool *radix.Pool
var ip_spam_script = radix.NewEvalScript(1,
	`local count = redis.call("INCR", KEYS[1])
redis.call("EXPIRE", KEYS[1], ARGV[1])
return count
`)

func init() {
	var err error
	pool, err = radix.NewPool("tcp", "localhost:6384", 200)
	if err != nil {
		Logger.Fatal("redis error", zap.Error(err))
	}

	Logger.Info("redis up", zap.String("port", ":6384"))
}

func main() {
	go func() {
		Logger.Info("prometheus", zap.String("port", ":9273"))
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9273", nil)
	}()

	Logger.Info("spam up", zap.Any("build", build))
	if err := run(); err != nil {
		Logger.Error("server error", zap.Error(err))
	}
}

func run() error {
	s := grpc.NewServer()
	spam_svc.RegisterSpamServer(s, &svc{})

	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		Logger.Error("listening error", zap.Error(err))
	}

	Logger.Info("listening", zap.String("port", ":3333"))
	return s.Serve(lis)
}
