package service

import (
	"github.com/mitchellh/mapstructure"
)

type Runner interface {
	SetContext(contextVariables []*Variable)
	SetOnDone(OnDone chan struct{})
	Done()
	Wait()
	GetId() string
	SetId(id string)
	GetError() error
	SetError(err error)
	GetOutputs() []*Variable
}
type Function struct {
	Id       string
	Name     string
	Outputs  []*Variable
	Context  []*Variable
	OnRun    func(f *Function, runtime *Runtime) error
	OnActive func(active *ActiveSignal, function *Function, runtime *Runtime)
	OnDone   chan struct{}
	Error    error
}

func (f *Function) SetError(err error) {
	f.Error = err
}

func (f *Function) GetError() error {
	return f.Error
}

func (f *Function) GetOutputs() []*Variable {
	return f.Outputs
}

func (f *Function) Wait() {
	<-f.OnDone
}

func (f *Function) Done() {
	f.OnDone <- struct{}{}
}

func (f *Function) SetId(id string) {
	f.Id = id
}

func (f *Function) GetId() string {
	return f.Id
}

func (f *Function) SetOnDone(OnDone chan struct{}) {
	f.OnDone = OnDone
}

func (f *Function) SetContext(contextVariables []*Variable) {
	f.Context = contextVariables
}

func (f *Function) getInputByName(name string) *Variable {
	for _, input := range f.Context {
		if input.Name == name {
			return input
		}
	}
	return nil
}
func DefaultOnActive(definition *FunctionDefinition, option BaseFunctionOption, active *ActiveSignal, function *Function) {
	for _, outputDef := range definition.Output {
		for _, output := range option.Outputs {
			if output.Name == outputDef.Name {
				for _, resultOutput := range active.Output {
					if resultOutput.Name == outputDef.Name {
						function.Outputs = append(function.Outputs, &Variable{
							Name:  output.Assign,
							Value: resultOutput.Value,
							Type:  resultOutput.Type,
						})
					}
				}
			}
		}
	}
}
func DefaultBlockOnActive(definition *FunctionDefinition, option BaseBlockOption, active *ActiveSignal, block *Block) {
	for _, outputDef := range definition.Output {
		for _, output := range option.Outputs {
			if output.Name == outputDef.Name {
				for _, resultOutput := range active.Output {
					if resultOutput.Name == outputDef.Name {
						block.Outputs = append(block.Outputs, &Variable{
							Name:  output.Assign,
							Value: resultOutput.Value,
							Type:  resultOutput.Type,
						})
					}
				}
			}
		}
	}
}

type Block struct {
	Id        string
	Outputs   []*Variable
	OutputDef []*FunctionOutput
	Context   []*Variable
	Runners   []Runner
	OnRun     func(block *Block, runtime *Runtime) error
	OnActive  func(active *ActiveSignal, function *Block, runtime *Runtime)
	OnDone    chan struct{}
	Error     error
}

func (b *Block) SetError(err error) {
	b.Error = err
}

func (b *Block) SetContext(contextVariables []*Variable) {
	b.Context = contextVariables
}

func (b *Block) SetOnDone(OnDone chan struct{}) {
	b.OnDone = OnDone
}

func (b *Block) Done() {
	b.OnDone <- struct{}{}
}

func (b *Block) Wait() {
	<-b.OnDone
}

func (b *Block) GetId() string {
	return b.Id
}

func (b *Block) SetId(id string) {
	b.Id = id
}

func (b *Block) GetError() error {
	return b.Error
}

func (b *Block) GetOutputs() []*Variable {
	return b.Outputs
}
func NewRootBlock(serviceContext ServiceContext) (*Block, error) {
	block := &Block{
		Outputs:   []*Variable{},
		Runners:   []Runner{},
		OutputDef: []*FunctionOutput{},
		OnRun: func(block *Block, runtime *Runtime) error {
			var err error
			for _, runner := range block.Runners {
				runner.SetOnDone(make(chan struct{}))
				runner.SetContext(block.Context)
				serviceContext.DefaultRuntime.Queue <- runner
				runner.Wait()
				block.Outputs = append(block.Outputs, runner.GetOutputs()...)
				err = runner.GetError()
				if err != nil {
					break
				}
			}
			serviceContext.DefaultRuntime.ActiveSignalChan <- &ActiveSignal{
				Id:     block.Id,
				Output: []*Variable{},
				Error:  err,
			}
			return nil
		},
		OnActive: func(active *ActiveSignal, block *Block, runtime *Runtime) {

		},
	}
	return block, nil
}
func NewBlock(def map[string]interface{}, serviceContext ServiceContext) (*Block, error) {
	var option BaseBlockOption
	err := mapstructure.Decode(def, &option)
	if err != nil {
		return nil, err
	}
	block := &Block{
		Outputs:   []*Variable{},
		Runners:   []Runner{},
		OutputDef: []*FunctionOutput{},
		OnRun: func(block *Block, runtime *Runtime) error {
			var err error
			for _, runner := range block.Runners {
				runner.SetOnDone(make(chan struct{}))
				runner.SetContext(block.Context)
				serviceContext.DefaultRuntime.Queue <- runner
				runner.Wait()
				block.Context = append(block.Context, runner.GetOutputs()...)
				err = runner.GetError()
				if err != nil {
					break
				}
			}
			serviceContext.DefaultRuntime.ActiveSignalChan <- &ActiveSignal{
				Id:     block.Id,
				Output: block.Context,
				Error:  err,
			}
			return nil
		},
		OnActive: func(active *ActiveSignal, block *Block, runtime *Runtime) {
			for _, outputVariableDef := range option.Outputs {
				for _, variable := range active.Output {
					if outputVariableDef.Name == variable.Name {
						block.Outputs = append(block.Outputs, &Variable{
							Name:  outputVariableDef.Assign,
							Value: variable.Value,
							Type:  variable.Type,
						})
					}
				}
			}
		},
	}
	return block, nil
}
