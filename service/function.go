package service

type Runner interface {
	Run()
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
func NewPlusFunction() *Function {
	return &Function{
		Name: "PlusFunction",
		OnRun: func(f *Function, runtime *Runtime) error {
			left := f.getInputByName("left")
			right := f.getInputByName("right")
			result := left.Value.(int) + right.Value.(int)
			runtime.ActiveSignalChan <- &ActiveSignal{Id: f.Id, Output: []*Variable{
				{
					Name:  "return",
					Value: result,
					Type:  "number",
				},
			}}
			return nil
		},
		Outputs: []*Variable{},
	}
}
func NewYouFileMoveFunction() *Function {
	return &Function{
		Name: "YouFileMoveFunction",
		OnRun: func(f *Function, runtime *Runtime) error {
			sourceVar := f.getInputByName("source")
			targetVar := f.getInputByName("target")
			DefaultYouFileClient.MoveFile([]*Variable{sourceVar, targetVar}, f)
			return nil
		},
		Outputs: []*Variable{},
	}
}
