package api

import (
	"github.com/guitarpawat/portscan/api/cache"
	"sync"
	"time"
)

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
	_, ok := t.Tasks[token]
	if !ok {
		t.Tasks[token] = task{
			Channels: []chan struct{}{},
		}
	}
}

func (t *tasks) registerChan(token string, kill chan struct{}) {
	tasks, ok := t.Tasks[token]
	if !ok {
		t.Tasks[token] = task{
			Channels: []chan struct{}{},
		}
	}
	tasks.Channels = append(tasks.Channels, kill)
}

func (t *tasks) revokeChan(token string) {
	m, ok := t.Tasks[token]
	if !ok {
		return
	}
	for _, v := range m.Channels {
		v <- struct{}{}
	}
	delete(t.Tasks, token)
}

func killTimeOut() {
	for k := range routines.Tasks {
		out, _ := cache.GetTokenInfo(k)
		if time.Since(out.LastUpdate) >= 3 * time.Minute {
			routines.revokeChan(k)
			cache.DeleteToken(k)
		}
	}
}

func registerToken(token string) {
	routines.Lock()
	defer routines.Unlock()
	routines.registerToken(token)
}

func registerChan(token string, kill chan struct{}) {
	routines.Lock()
	defer routines.Unlock()
	routines.registerChan(token, kill)
}

func revokeChan(token string)  {
	routines.Lock()
	defer routines.Unlock()
	routines.revokeChan(token)
}