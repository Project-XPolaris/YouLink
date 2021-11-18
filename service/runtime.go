package service

import (
	"context"
	"github.com/rs/xid"
	"sync"
)

type ActiveSignal struct {
	Id     string
	Output []*Variable
	Error  error
}
type Runtime struct {
	sync.Mutex
	Queue            chan Runner
	Current          Runner
	Suspend          []Runner
	ActiveSignalChan chan *ActiveSignal
}

func NewRuntime() *Runtime {
	return &Runtime{
		Queue:            make(chan Runner, 1000),
		ActiveSignalChan: make(chan *ActiveSignal, 1000),
		Suspend:          []Runner{},
	}
}
func (r *Runtime) GetSuspendRunner(id string) Runner {
	for _, function := range r.Suspend {
		if id == function.GetId() {
			return function
		}
	}
	return nil
}
func (r *Runtime) RemoveSuspendRunner(id string) Runner {
	r.Lock()
	defer r.Unlock()
	newList := make([]Runner, 0)
	for _, function := range r.Suspend {
		if id != function.GetId() {
			newList = append(newList, function)
		}
	}
	r.Suspend = newList
	return nil
}
func (r *Runtime) Run(ctx context.Context) {
	for {
		select {
		case runner := <-r.Queue:
			r.Lock()
			runner.SetId(xid.New().String())
			r.Suspend = append(r.Suspend, runner)
			r.Unlock()
			go func() {
				switch runner.(type) {
				case *Function:
					function := runner.(*Function)
					err := function.OnRun(function, r)
					if err != nil {
						r.ActiveSignalChan <- &ActiveSignal{
							Id:     function.Id,
							Output: function.Outputs,
							Error:  err,
						}
					}
				case *Block:
					block := runner.(*Block)
					err := block.OnRun(block, r)
					if err != nil {
						r.ActiveSignalChan <- &ActiveSignal{
							Id:     block.Id,
							Output: block.Outputs,
							Error:  err,
						}
					}
				}

			}()
		case active := <-r.ActiveSignalChan:
			runner := r.GetSuspendRunner(active.Id)
			if active.Error != nil {
				runner.SetError(active.Error)
			}
			switch runner.(type) {
			case *Function:
				function := runner.(*Function)
				if function.OnActive != nil {
					function.OnActive(active, function, r)
				}
			case *Block:
				block := runner.(*Block)
				if block.OnActive != nil {
					block.OnActive(active, block, r)
				}
			}
			r.RemoveSuspendRunner(active.Id)
			runner.Done()
		case <-ctx.Done():
			return
		}
	}
}
