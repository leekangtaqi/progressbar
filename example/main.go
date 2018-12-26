package main

import (
	"testing"
	"time"

	"github.com/leekangtaqi/progressbar/v1"
)

func TestProgressBar(t *testing.T) {
	setPg, _ := progressbar.New()
	setPg(0.1256)
	time.Sleep(time.Millisecond * 100)
	setPg(0.2244)
	time.Sleep(time.Millisecond * 100)
	setPg(0.3251)
	time.Sleep(time.Millisecond * 100)
	setPg(0.4221)
	time.Sleep(time.Millisecond * 100)
	setPg(0.4872)
	time.Sleep(time.Millisecond * 100)
	setPg(0.5821212)
	time.Sleep(time.Millisecond * 100)
	setPg(0.618388)
	time.Sleep(time.Millisecond * 100)
	setPg(0.778844)
	time.Sleep(time.Millisecond * 100)
	setPg(0.91877)
	time.Sleep(time.Millisecond * 100)
	setPg(0.92877)
	time.Sleep(time.Millisecond * 2000)
	setPg(1)
	time.Sleep(time.Millisecond * 200)
}
