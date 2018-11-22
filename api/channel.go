package api

import "sync"

type tasks struct {
	Tasks map[string]task
	sync.Mutex
}

type task struct {
	Channels []chan struct{}
}

var routines = &tasks{
	Tasks: map[string]task{},
	Mutex: sync.Mutex{},
}

func (t *tasks) registerToken(token string) {
	t.Lock()
	defer t.Unlock()
	_, ok := t.Tasks[token]
	if !ok {
		t.Tasks[token] = task{
			Channels: []chan struct{}{},
		}
	}
}

func (t *tasks) registerChan(token string, kill chan struct{}) {
	t.Lock()
	defer t.Unlock()
	tasks, ok := t.Tasks[token]
	if !ok {
		t.Tasks[token] = task{
			Channels: []chan struct{}{},
		}
	}
	tasks.Channels = append(tasks.Channels, kill)
}

func (t *tasks) revokeChan(token string) {
	t.Lock()
	defer t.Unlock()
	m, ok := t.Tasks[token]
	if !ok {
		return
	}
	for _, v := range m.Channels {
		v <- struct{}{}
	}
	delete(t.Tasks, token)
}