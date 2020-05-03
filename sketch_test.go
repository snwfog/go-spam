package main

import (
	"bytes"
	"math"
	"math/rand"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/shenwei356/countminsketch"
	"github.com/stretchr/testify/assert"
)

func Test_sketch_copy_restore(t *testing.T) {
	expect := assert.New(t)

	key := []byte(`123`)
	sketch, _ := countminsketch.NewWithEstimates(0.1, 0.999)
	sketch.Update(key, 1)

	expect.Equal(uint64(1), sketch.Estimate(key))

	var buf bytes.Buffer
	_, err := sketch.WriteTo(&buf)
	expect.NoError(err)

	dest, _ := countminsketch.NewWithEstimates(0.1, 0.999)
	_, err = dest.ReadFrom(bytes.NewReader(buf.Bytes()))
	expect.NoError(err)

	expect.Equal(uint64(1), sketch.Estimate(key))
}

func Test_tsk(t *testing.T) {
	expect := assert.New(t)
	sk, err := create()

	expect.NoError(err)

	key := []byte(`123`)
	est, err := sk.Estimate(key, 1)
	expect.NoError(err)
	expect.Equal(uint64(1), est)

	est, err = sk.Estimate(key, 1)
	expect.NoError(err)
	expect.Equal(uint64(2), est)
}

// go test -run=sketch -v
// === RUN   Test_sketch
//    Test_sketch: hist.go:41: k:                    0, v:      69005
//    Test_sketch: hist.go:41: k:                    1, v:     263185
//    Test_sketch: hist.go:41: k:                    2, v:     369928
//    Test_sketch: hist.go:41: k:                    3, v:     216508
//    Test_sketch: hist.go:41: k:                    4, v:      53570
//    Test_sketch: hist.go:41: k:                    5, v:       5742
//    Test_sketch: hist.go:41: k:                    6, v:        353
//    Test_sketch: hist.go:41: k:                    7, v:         12
// --- PASS: Test_sketch (6.17s)
// PASS
// ok  	github.com/StackAdapt/go-spam	6.258s
func Test_sketch(t *testing.T) {
	expect := assert.New(t)
	sk, _ := countminsketch.NewWithEstimates(0.00001, 0.999)

	thr := 1 << 10
	var wg sync.WaitGroup
	var m sync.Map

	wg.Add(thr)

	for i := 0; i < thr; i++ {
		go func() {
			for j := 0; j < 1<<10; j++ {
				ip := rand_ip()
				ipp := net.IPv4(ip[0], ip[1], ip[2], ip[3])
				// t.Logf("ip: %s", ipp)

				prev, ok := m.Load(ipp.String())
				if !ok {
					prev = interface{}(0)
				}

				sk.Update(ip, 1)
				m.Store(ipp.String(), prev.(int)+1)
			}

			wg.Done()
		}()
	}

	h := new_hist()
	wg.Wait()
	m.Range(func(k, v interface{}) bool {
		ip := net.ParseIP(k.(string))
		if ip == nil {
			t.Errorf("wrong ip, %s", k)
			return true
		}

		hit_count, sk_estimate := v.(int), sk.Estimate(ip.To4())
		// t.Logf("ip: %20s, hit: %d, estimate: %d", k, v, sk.Estimate(ip.To4()))
		h.Add(int64(math.Abs(float64(hit_count-int(sk_estimate)))), 1)
		return true
	})

	expect.True(true)
	h.Print(t)
}

var rand_pool = sync.Pool{
	New: func() interface{} {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

func rand_ip() []byte {
	var ip [4]byte
	rand_pool.Get().(*rand.Rand).Read(ip[:])
	return ip[:]
}
