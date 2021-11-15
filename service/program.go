package service

import "github.com/rs/xid"

type Program struct {
	Id        string
	Name      string
	Context   []*Variable
	Functions []*Function
	OnDone    chan struct{}
	Error     error
}

func NewProgram() *Program {
	return &Program{
		Id:        xid.New().String(),
		Context:   []*Variable{},
		Functions: []*Function{},
		OnDone:    make(chan struct{}, 0),
	}
}

func (p *Program) Run(runtime *Runtime) error {
	for _, function := range p.Functions {
		function.OnDone = make(chan struct{}, 0)
		function.Context = p.Context
		runtime.Queue <- function
		<-function.OnDone
		for _, output := range function.Outputs {
			p.Context = append(p.Context, output)
		}
		if function.Error != nil {
			p.Error = function.Error
			return function.Error
		}
	}
	return nil
}
