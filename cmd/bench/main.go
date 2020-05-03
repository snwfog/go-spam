package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	grpcpool "github.com/snwfog/grpc-go-pool"
	"google.golang.org/grpc"

	spam_svc "github.com/StackAdapt/go-spam/external/services/grpc/definition"
)

var conn_pool *grpcpool.Pool
var rand_pool sync.Pool

func init() {
	var err error
	conn := func() (*grpc.ClientConn, error) {
		return grpc.Dial("localhost:3333",
			grpc.WithInsecure(),
			grpc.WithDefaultCallOptions())
	}
	conn_pool, err = grpcpool.New(conn, 10, 100, 120*time.Second, 120*time.Second)
	if err != nil {
		log.Fatalf("error creating connection pool: %v", err)
	}

	// Pool of prng
	rand_pool = sync.Pool{
		New: func() interface{} {
			return rand.New(rand.NewSource(time.Now().UnixNano()))
		},
	}
}

func main() {
	thr := 1 << 10

	var wg sync.WaitGroup
	wg.Add(thr)

	for i := 0; i < thr; i++ {
		go func() {
			for j := 0; j < 1<<10; j++ {
				ip := rand_ip()
				// ip := []byte{8, 8, 8, 8}
				log.Printf("ip: %20s, estimate: %d\n", net.IPv4(ip[0], ip[1], ip[2], ip[3]), make_call(ip))
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func make_call(ip []byte) uint64 {
	pool_ctx, pool_cancel := get_ctx(500 * time.Millisecond)
	cc, err := conn_pool.Get(pool_ctx)
	if err != nil {
		log.Println("error getting connection from pool", err)
	}

	pool_cancel()
	defer cc.Close()

	ctx, cancel := get_ctx(1 * time.Second)
	defer cancel()

	client := spam_svc.NewSpamClient(cc.ClientConn)
	resp, err := client.SetGet(ctx, &spam_svc.Request{Key: ip, Count: 1})

	if err != nil {
		log.Println("error: bid", err)
		return 0
	}

	// log.Println("hit count:", resp.Count)
	return resp.Count
}

func get_ctx(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}

func rand_ip() []byte {
	var ip [4]byte
	rand_pool.Get().(*rand.Rand).Read(ip[:])
	return ip[:]
}
