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
		instance = &watcher{watchers: make(map[int]*state)}
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

func (w *watcher) GetOpenTime(term int) (op time.Time, err error) {
	if b, e := w.IsOpen(term); e != nil {
		return time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), errors.New("terminal not opened")
	} else {
		if !b {
			err = errors.New("terminal already close")
		} else {
			err = nil
		}
		op = w.watchers[term].openTime
		return
	}
}

func (w *watcher) CloseTerminal(term int) {
	w.watchers[term].isOpen = false
}
