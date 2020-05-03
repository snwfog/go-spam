package main

import (
	"sync"

	"github.com/gogo/protobuf/sortkeys"
)

type hist struct {
	m map[int64]uint64

	sync.RWMutex
}

func new_hist() *hist {
	h := hist{}
	h.m = make(map[int64]uint64)

	return &h
}

func (h *hist) Add(key int64, count uint64) {
	h.Lock()
	defer h.Unlock()

	h.m[key] = h.m[key] + count
}

type ILog interface {
	Logf(string, ...interface{})
}

func (h *hist) Print(printer ILog) {
	keys := make([]int64, 0, len(h.m))
	for count := range h.m {
		keys = append(keys, count)
	}

	sortkeys.Int64s(keys)
	for _, count := range keys {
		printer.Logf("k: %20d, v: %10d", count, h.m[count])
	}
}
