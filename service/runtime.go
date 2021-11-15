package service

import (
	"context"
	"github.com/rs/xid"
	"sync"
)

var DefaultRuntime = NewRuntime()

type ActiveSignal struct {
	Id     string
	Output []*Variable
	Error  error
}
type Runtime struct {
	sync.Mutex
	Queue            chan *Function
	Current          *Function
	Suspend          []*Function
	ActiveSignalChan chan *ActiveSignal
}

func NewRuntime() *Runtime {
	return &Runtime{
		Queue:            make(chan *Function, 1000),
		ActiveSignalChan: make(chan *ActiveSignal, 1000),
		Suspend:          []*Function{},
	}
}
func (r *Runtime) GetSuspendFunction(id string) *Function {
	for _, function := range r.Suspend {
		if id == function.Id {
			return function
		}
	}
	return nil
}
func (r *Runtime) RemoveSuspendFunction(id string) *Function {
	r.Lock()
	defer r.Unlock()
	newList := make([]*Function, 0)
	for _, function := range r.Suspend {
		if id != function.Id {
			newList = append(newList, function)
		}
	}
	r.Suspend = newList
	return nil
}
func (r *Runtime) Run(ctx context.Context) {
	for {
		select {
		case function := <-r.Queue:
			r.Lock()
			function.Id = xid.New().String()
			r.Suspend = append(r.Suspend, function)
			r.Unlock()
			go func() {
				function.OnRun(function, r)
			}()
		case active := <-r.ActiveSignalChan:
			function := r.GetSuspendFunction(active.Id)
			if function.OnActive != nil {
				function.OnActive(active, function, r)
			}
			r.RemoveSuspendFunction(active.Id)
			function.OnDone <- struct{}{}
		case <-ctx.Done():
			return
		}
	}
}
