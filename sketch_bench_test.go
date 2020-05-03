package main

import (
	"testing"
)

var Count uint64

// go test -run=XXX -bench=tsk -benchtime=30s -benchmem -v
// 2020-05-02T23:04:37.307-0400	INFO	redis up	{"port": ":6384"}
// goos: darwin
// goarch: amd64
// pkg: github.com/StackAdapt/go-spam
// Benchmark_tsk
// Benchmark_tsk-16    	170908729	       216 ns/op	     294 B/op	       6 allocs/op
// PASS
// ok  	github.com/StackAdapt/go-spam	58.197s
func Benchmark_tsk(b *testing.B) {
	ip := rand_ip()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Count, _ = Estimate(ip, 1)
		}
	})
}
