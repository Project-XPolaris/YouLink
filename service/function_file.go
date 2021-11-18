package service

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"path/filepath"
)

type BaseFunctionOption struct {
	Inputs  []FunctionInput
	Outputs []FunctionOutput
}

func (o *BaseFunctionOption) GetInputByName(name string) *FunctionInput {
	for _, input := range o.Inputs {
		if input.Name == name {
			return &input
		}
	}
	return nil
}
func (o *BaseFunctionOption) CheckValidate(entity FunctionEntity) error {
	definition := entity.GetDefinition()
	for _, variableDefinition := range definition.Input {
		inputVar := o.GetInputByName(variableDefinition.Name)
		if inputVar == nil {
			return errors.New(fmt.Sprintf("invalidate function [%s] input:[%s] not found", definition.Name, variableDefinition.Name))
		}
	}
	return nil
}

type BaseBlockOption struct {
	Outputs []FunctionOutput
}
type UpdatePathDirFunction struct {
	Definition *FunctionDefinition
}

func (f *UpdatePathDirFunction) GetDefinition() *FunctionDefinition {
	return f.Definition
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
			sourceVar, err := ReadVariable(function, "source", "string", &option)
			if err != nil {
				return err
			}
			targetVar, err := ReadVariable(function, "target", "number", &option)
			if err != nil {
				return err
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
