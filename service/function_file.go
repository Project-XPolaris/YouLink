package service

import (
	"github.com/mitchellh/mapstructure"
	"path/filepath"
)

type BaseFunctionOption struct {
	Inputs  []FunctionInput
	Outputs []FunctionOutput
}
type UpdatePathDirFunction struct {
	Definition *FunctionDefinition
}

func NewUpdatePathDirFunction() *UpdatePathDirFunction {
	return &UpdatePathDirFunction{
		Definition: &FunctionDefinition{
			Name: "UpdateDirectory",
			Input: []*VariableDefinition{
				{
					Name: "source",
					Type: "string",
				},
				{
					Name: "target",
					Type: "string",
				},
			},
			Output: []*VariableDefinition{
				{
					Name: "path",
					Type: "string",
				},
			},
		},
	}

}
func (f *UpdatePathDirFunction) ToFunction(m map[string]interface{}) (*Function, error) {
	var option BaseFunctionOption
	err := mapstructure.Decode(m, &option)
	if err != nil {
		return nil, err
	}

	function := &Function{
		Name: "UpdateFileDir",
		OnRun: func(function *Function, runtime *Runtime) error {
			var sourceVar, targetVar *Variable
			for _, input := range option.Inputs {
				if input.Name == "source" {
					if len(input.Ref) > 0 {
						sourceVar = function.getInputByName(input.Ref)
					} else {
						sourceVar = &Variable{
							Name:  "source",
							Value: input.Value.(string),
							Type:  "string",
						}
					}
				}
				if input.Name == "target" {
					if len(input.Ref) > 0 {
						targetVar = function.getInputByName("ref")
					} else {
						targetVar = &Variable{
							Name:  "target",
							Value: input.Value.(string),
							Type:  "string",
						}
					}
				}
			}
			base := filepath.Base(sourceVar.Value.(string))
			result := filepath.Join(targetVar.Value.(string), base)
			outputVars := make([]*Variable, 0)
			outputVars = append(outputVars, &Variable{
				Name:  "path",
				Value: result,
				Type:  "string",
			})
			runtime.ActiveSignalChan <- &ActiveSignal{
				Id:     function.Id,
				Output: outputVars,
				Error:  nil,
			}
			return nil
		},
		OnActive: func(active *ActiveSignal, function *Function, runtime *Runtime) {
			DefaultOnActive(f.Definition, option, active, function)
		},
	}
	return function, nil
}
