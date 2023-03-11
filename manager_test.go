package owm

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestMananger(t *testing.T) {
	var m = NewManager(5)
	m.Submit(1, func(w *Worker) {
		fmt.Println("在Worker", w.id, "中执行了")
	})
	m.Submit(2, func(w *Worker) {
		fmt.Println("在Worker", w.id, "中执行了")
	})
	m.Submit(3, func(w *Worker) {
		time.Sleep(time.Second * 2)
		fmt.Println("在Worker", w.id, "中延迟执行了")
	})
	m.Stop()
}

func BenchmarkWorkerSubmitQPS(b *testing.B) {
	var m = NewManager(runtime.NumCPU())
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Submit(i, func(w *Worker) {})
	}
	m.Stop()
}
