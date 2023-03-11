package owm

import (
	"context"
)

type Manager struct {
	workers []*Worker
	ctx     context.Context
}

type task func(w *Worker)

func NewManager(n int) *Manager {
	var m = &Manager{
		workers: make([]*Worker, n),
		ctx:     context.Background(),
	}
	for i := 0; i < len(m.workers); i++ {
		m.workers[i] = NewWorker(i, m.ctx)
		m.workers[i].Start()
	}
	return m
}

func (m *Manager) Submit(workerId int, t task) {
	if workerId >= len(m.workers) {
		workerId %= len(m.workers)
	}
	m.workers[workerId].Submit(t)
}

func (m *Manager) Stop() {
	for _, worker := range m.workers {
		worker.stop()
	}
}
