package service

import (
	"context"
	"sync"
)

type ProgramLauncher struct {
	sync.Mutex
	Suspend []*Program
	Queue   chan *Program
}

func NewProgramLauncher() *ProgramLauncher {
	return &ProgramLauncher{
		Suspend: []*Program{},
		Queue:   make(chan *Program, 1000),
	}
}
func (l *ProgramLauncher) RemoveSuspendProgram(id string) {
	l.Lock()
	defer l.Unlock()
	var newList []*Program
	for _, program := range l.Suspend {
		if program.Id != id {
			newList = append(newList, program)
		}
	}
	l.Suspend = newList
}
func (l *ProgramLauncher) Run(context context.Context, serviceContext *ServiceContext) {
	for {
		select {
		case program := <-l.Queue:
			l.Suspend = append(l.Suspend, program)
			go func() {
				program.Run(serviceContext.DefaultRuntime)
				l.RemoveSuspendProgram(program.Id)
				close(program.OnDone)
			}()
		case <-context.Done():
			return
		}
	}
}
