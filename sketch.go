package main

import (
	"bytes"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/shenwei356/countminsketch"
	"github.com/valyala/bytebufferpool"
	uatomic "go.uber.org/atomic"
	"go.uber.org/zap"

	. "github.com/StackAdapt/go-spam/log"
)

const epsilon = 0.00001
const delta = 0.9999

const w_size = 7      // maximum window size
const b_size = 1 << 9 // bucket size
const t_max = 5       // maximum duration, allow 2 seconds drift

var time_series_sketch *tsk

func init() {
	s, err := create()
	if err != nil {
		Logger.Fatal("error init sk", zap.Error(err))
	}

	go s.run()

	time_series_sketch = s
}

func Estimate(key []byte, count uint64) (uint64, error) {
	return time_series_sketch.Estimate(key, count)
}

// region Time Series Sketch

// TODO: Add head and tail ptr?
type tsk struct {
	*bucket
}

type bucket struct {
	g_clock *uatomic.Int64

	w [w_size]unsafe.Pointer

	mu sync.RWMutex
}

func (b *bucket) run() {
	for range time.Tick(time.Second) {
		// b.mu.Lock()

		g_time := b.g_clock.Load()
		// Logger.Debug("updating g_clock", zap.Int64("g_time", g_time))

		buf := bytebufferpool.Get()

		// copy from g_clock to g_clock+1
		src_sk := (*countminsketch.CountMinSketch)(atomic.LoadPointer(&(b.w[g_time%w_size])))
		// dest_sk := (*countminsketch.CountMinSketch)(atomic.LoadPointer(&b.w[(g_time+1)%w_size]))
		dest_sk, err := countminsketch.NewWithEstimates(epsilon, delta)
		if err != nil {
			Logger.Fatal("error new", zap.Error(err))
		}

		// Logger.Debug("clock estimate",
		// 	zap.Uint64("old", src_sk.Estimate(g_key)))

		if _, err := src_sk.WriteTo(buf); err != nil {
			Logger.Fatal("error ser", zap.Error(err))
		}

		if _, err := dest_sk.ReadFrom(bytes.NewReader(buf.Bytes())); err != nil {
			Logger.Fatal("error deser", zap.Error(err))
		}

		atomic.StorePointer(&(b.w[(g_time+1)%w_size]), unsafe.Pointer(dest_sk))
		b.g_clock.Store(g_time + 1)
		// Logger.Info("g_clock",
		// 	zap.Int64("start time", g_time), zap.Int64("g_time", b.g_clock.Load()))

		// b.mu.Unlock()
	}
}

func create() (*tsk, error) {
	var sk tsk

	bucket := bucket{g_clock: uatomic.NewInt64(time.Now().Unix())}
	err := new_w(bucket.w[:])
	if err != nil {
		Logger.Fatal("error creating bucket", zap.Error(err))
	}

	sk = tsk{
		bucket: &bucket,
	}
	return &sk, nil
}

func (t *tsk) run() {
	t.bucket.run()
}

// TODO: Handle extremely large count (negatived), or should let go?
func (t *tsk) Estimate(key []byte, count uint64) (uint64, error) {
	now, g_time := time.Now().Unix(), t.g_clock.Load()
	if now <= g_time {
		// good scenario
		return t.estimate(key, count, 5)
	}

	// Logger.Debug("now > g_clock (forced g_time)",
	// 	zap.Int64("now", now), zap.Int64("g_time", g_time))
	return t.estimate(key, count, 5)
}

func (t *tsk) estimate(key []byte, count uint64, duration uint64) (uint64, error) {
	// t.bucket.mu.Lock()
	// defer t.bucket.mu.Unlock()

	end := uint64(t.g_clock.Load())

	start_sk := (*countminsketch.CountMinSketch)(atomic.LoadPointer(&(t.w[(end-duration)%w_size])))
	now_sk := (*countminsketch.CountMinSketch)(atomic.LoadPointer(&(t.w[end%w_size])))

	now_sk.Update(key, count)
	count_start, count_now := start_sk.Estimate(key), now_sk.Estimate(key)
	// Logger.Debug("query estimate", zap.ByteString("key", key),
	// 	zap.String("range", fmt.Sprintf("[%d-%d]", end-duration, end)),
	// 	zap.Uint64("start", count_start),
	// 	zap.Uint64("now", count_now))
	return count_now - count_start, nil
}

func new_w(w []unsafe.Pointer) error {
	for j := 0; j < w_size; j++ {
		sketch, err := countminsketch.NewWithEstimates(epsilon, delta)
		if err != nil {
			return err
		}

		atomic.StorePointer(&(w[j]), unsafe.Pointer(sketch))
	}

	return nil
}

// endregion
