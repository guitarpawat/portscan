package api

import "sync"

type tasks struct {
	Tasks []task
	sync.Mutex
}

type task struct {
	Token    string
	Channels []chan struct{}
}

var routines = &tasks{
	Tasks: []task{},
	Mutex: sync.Mutex{},
}