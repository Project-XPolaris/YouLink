package service

import (
	"errors"
	"github.com/rs/xid"
	"sync"
)

var (
	ProcessStatusRunning  = "Running"
	ProcessStatusEstimate = "Estimate"
	ProcessStatusError    = "Error"
	ProcessStatusDone     = "Done"
)

type ProcessManager struct {
	sync.Mutex
	Processes []*Process
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		Processes: []*Process{},
	}
}

func (m *ProcessManager) AllocateProcess() *Process {
	process := NewProcess()
	m.Lock()
	defer m.Unlock()
	m.Processes = append(m.Processes, process)
	return process
}
func (m *ProcessManager) GetProcessById(id string) *Process {
	for _, process := range m.Processes {
		if process.Id == id {
			return process
		}
	}
	return nil
}

type Process struct {
	Id       string
	Programs []*Program
	Status   string
	Err      error
}

func NewProcess() *Process {
	return &Process{
		Id:       xid.New().String(),
		Programs: []*Program{},
		Status:   ProcessStatusEstimate,
	}
}

func (p *Process) AddProgram(program *Program) error {
	if p.Status != ProcessStatusEstimate {
		return errors.New("must in estimate status")
	}
	p.Programs = append(p.Programs, program)
	return nil
}

func (p *Process) SetInput(variables []*Variable) error {
	if p.Status != ProcessStatusEstimate {
		return errors.New("must in estimate status")
	}
	if len(p.Programs) == 0 {
		return nil
	}
	for _, program := range p.Programs {
		program.Context = append(p.Programs[0].Context, variables...)
	}
	return nil
}
func (p *Process) Run(ctx ServiceContext) {
	p.Status = ProcessStatusRunning
	go func() {
		for _, program := range p.Programs {
			err := program.Run(ctx.DefaultRuntime)
			if err != nil {
				p.Status = ProcessStatusError
				p.Err = err
				return
			}
		}
		p.Status = ProcessStatusDone
	}()
}
