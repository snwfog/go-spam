package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"
	"time"

	grpcpool "github.com/snwfog/grpc-go-pool"
	"google.golang.org/grpc"

	spam_svc "github.com/StackAdapt/go-spam/external/services/grpc/definition"
)

var connpool *grpcpool.Pool

func init() {
	var err error
	conn := func() (*grpc.ClientConn, error) {
		// https://github.com/bradleyjkemp/grpc-tools/issues/29
		// https://weblogs.asp.net/owscott/introducing-testing-domain-localtest-me
		// return grpc.Dial("localtest.me:7700",
		// 	grpc.WithInsecure(),
		// 	grpc.WithDefaultCallOptions(
		// 		grpc.ForceCodec(&codec.Msgpack{})))

		return grpc.Dial("localhost:3333",
			grpc.WithInsecure(),
			grpc.WithDefaultCallOptions())
	}
	connpool, err = grpcpool.New(conn, 10, 100, 120*time.Second, 120*time.Second)
	if err != nil {
		log.Fatalf("error creating connection pool: %v", err)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() <= 0 {
		log.Fatalf(`
Usage: go run main.go <ip>
Usage: go run main.go <ip> <count>
`)
	}

	key := flag.Arg(0)
	count, _ := strconv.Atoi(flag.Arg(1))
	log.Printf("ip: %s, count: %s", key, count)
	log.Println()

	ip := net.ParseIP(key)
	if ip == nil {
		log.Fatalf("invalid ip %s", key)
	}

	cc, err := connpool.Get(context.Background())
	if err != nil {
		log.Fatalf("error getting connection from pool %+v", err)
	}
	defer cc.Close()

	ctx, cancel := getctx(time.Second)
	defer cancel()

	client := spam_svc.NewSpamClient(cc.ClientConn)
	resp, err := client.SetGet(ctx, &spam_svc.Request{Key: ip.To4(), Count: uint64(count)})

	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	log.Printf("hit count: %d", resp.Count)
}

func getctx(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}
