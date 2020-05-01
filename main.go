package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	spam_svc "github.com/StackAdapt/go-spam/external/services/grpc/definition"
)

func main() {
	logrus.Infof("spam up: %+v", build)
	if err:= run(); err != nil {
		logrus.Errorf("server error: %+v", err)
	}
}

func run() error {
	s := grpc.NewServer()
	spam_svc.RegisterSpamServer(s, &svc{})

	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		logrus.Error("listening error", err)
	}

	logrus.WithFields(logrus.Fields{"port": ":3333"}).Info("listening")
	return s.Serve(lis)
}

type svc struct{}

func (svc) SetGet(ctx context.Context, req *spam_svc.Request) (*spam_svc.Response, error) {
	estimate, err := Estimate(req.Key, req.Count)
	if err != nil {
		return nil, fmt.Errorf("error: %+v", err)
	}

	return &spam_svc.Response{Count: estimate}, nil
}
