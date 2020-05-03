package main

import (
	"testing"
)

var Count uint64
func Benchmark_tsk(b *testing.B) {
	ip := rand_ip()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Count, _ = Estimate(ip, 1)
		}
	})
}
