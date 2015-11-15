package utils

import (
	"errors"
	"sync"
	"time"
)

type watcher struct {
	watchers map[int]*state
}

type state struct {
	isOpen   bool
	openTime time.Time
}

var instance *watcher
var once sync.Once

func GetWatcherInstance() *watcher {
	once.Do(func() {
		instance = &watcher{}
	})

	return instance
}

func (w *watcher) OpenTerminal(term int, aTime time.Time) {
	w.watchers[term] = &state{true, aTime}
}

func (w *watcher) IsOpen(term int) (bool, error) {
	val, ok := w.watchers[term]
	if ok {
		return val.isOpen, nil
	} else {
		return false, errors.New("terminal not found")
	}
}

func (w *watcher) CloseTerminal(term int) {
	w.watchers[term].isOpen = false
}
