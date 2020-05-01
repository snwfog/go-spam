package main

import (
	"bytes"
	"time"

	"github.com/shenwei356/countminsketch"
	"github.com/sirupsen/logrus"
	"github.com/valyala/bytebufferpool"
	"go.uber.org/atomic"
)

var time_series_sketch *tsk
var g_time *atomic.Int64

func init() {
	s, err := create()
	if err != nil {
		logrus.Fatal("error init sk")
	}

	time_series_sketch = s
	g_time = atomic.NewInt64(time.Now().Unix())
}

func Estimate(key []byte, count uint64) (uint64, error) {
	return time_series_sketch.Estimate(key, count, 5)
}

// region Time Series Sketch
type tsk struct {
	w [7]*countminsketch.CountMinSketch
}

func create() (*tsk, error) {
	var sk tsk

	for i := 0; i < 7; i++ {
		sketch, err := countminsketch.NewWithEstimates(0.1, 0.999)
		if err != nil {
			return nil, err
		}

		sk.w[i] = sketch
	}

	return &sk, nil
}

func (t *tsk) Estimate(key []byte, count uint64, duration uint8) (uint64, error) {
	now, g_clock := time.Now().Unix(), g_time.Load()
	if now <= g_clock { // if time is less than global
		return t.estimate(key, count, duration, now)
	}

	// else we need update
	if g_time.CAS(g_clock, now) {
		buf := bytebufferpool.Get()
		if _, err := t.w[g_clock%7].WriteTo(buf); err != nil {
			logrus.Fatal("error ser", err)
		}

		if _, err := t.w[now%7].ReadFrom(bytes.NewReader(buf.Bytes())); err != nil {
			logrus.Fatal("error deser", err)
		}

		now = g_clock
	}

	return t.estimate(key, count, duration, now)
}

func (t *tsk) estimate(key []byte, count uint64, duration uint8, end int64) (uint64, error) {
	start, now := t.w[(end-int64(duration))%7], t.w[end%7]
	now.Update(key, count)

	count_start, count_end := start.Estimate(key), now.Estimate(key)
	return count_end - count_start, nil
}

// endregion
