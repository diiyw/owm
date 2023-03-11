package owm

import (
	"context"
	"log"
)

type Worker struct {
	id     int
	ctx    context.Context
	cancel context.CancelFunc
	task   chan task
	done   chan struct{}
}

func NewWorker(id int, ctx context.Context) *Worker {
	w := &Worker{
		id:   id,
		task: make(chan task, 8),
		done: make(chan struct{}),
	}
	w.ctx, w.cancel = context.WithCancel(ctx)
	return w
}

func (w *Worker) Submit(t task) {
	select {
	case <-w.ctx.Done():
		return
	default:
		w.task <- t
	}
}

func (w *Worker) stop() {
	w.cancel()
	<-w.done
}

func (w *Worker) Start() {
	go func() {
		log.Println("Worker:", w.id, " running. ")
		for {
			select {
			case <-w.ctx.Done():
				if len(w.task) == 0 {
					log.Println("Worker:", w.id, " closed. ")
					w.done <- struct{}{}
					return
				}
			case task := <-w.task:
				if task != nil {
					task(w)
				}
			}
		}
	}()
}
