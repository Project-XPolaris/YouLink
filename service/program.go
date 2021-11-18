package service

import "github.com/rs/xid"

type Program struct {
	Id      string
	Name    string
	Context []*Variable
	Runners []Runner
	OnDone  chan struct{}
	Error   error
}

func NewProgram() *Program {
	return &Program{
		Id:      xid.New().String(),
		Context: []*Variable{},
		Runners: []Runner{},
		OnDone:  make(chan struct{}, 0),
	}
}

func (p *Program) Run(runtime *Runtime) error {
	for _, runner := range p.Runners {
		runner.SetOnDone(make(chan struct{}, 0))
		runner.SetContext(p.Context)
		runtime.Queue <- runner
		runner.Wait()
		for _, output := range runner.GetOutputs() {
			newVariable := true
			for _, variable := range p.Context {
				if variable.Name == output.Name {
					variable.Value = output.Value
					newVariable = false
					break
				}
			}
			if newVariable {
				p.Context = append(p.Context, output)
			}
		}
		if runner.GetError() != nil {
			p.Error = runner.GetError()
			return runner.GetError()
		}
	}
	return nil
}
