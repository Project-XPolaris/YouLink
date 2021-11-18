package service

import (
	"github.com/mitchellh/mapstructure"
	"strings"
)

type ConcatFunctionEntity struct {
}

func NewConcatFunctionEntity() *ConcatFunctionEntity {
	return &ConcatFunctionEntity{}
}

func (p *ConcatFunctionEntity) GetName() string {
	return "Concat"
}
func (p *ConcatFunctionEntity) ToFunction(m map[string]interface{}) (*Function, error) {
	var option BaseFunctionOption
	err := mapstructure.Decode(m, &option)
	if err != nil {
		return nil, err
	}
	err = option.CheckValidate(p)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	function := &Function{
		Name: p.GetName(),
		OnRun: func(f *Function, runtime *Runtime) error {
			left, err := ReadVariable(f, "left", "string", &option)
			if err != nil {
				return err
			}
			right, err := ReadVariable(f, "right", "string", &option)
			if err != nil {
				return err
			}
			result := left.Value.(string) + right.Value.(string)
			runtime.ActiveSignalChan <- &ActiveSignal{Id: f.Id, Output: []*Variable{
				{
					Name:  "return",
					Value: result,
					Type:  "string",
				},
			}}
			return nil
		},
		Outputs: []*Variable{},
		OnActive: func(active *ActiveSignal, function *Function, runtime *Runtime) {
			DefaultOnActive(p.GetDefinition(), option, active, function)
		},
	}
	return function, nil

}

func (p *ConcatFunctionEntity) GetDefinition() *FunctionDefinition {
	return &FunctionDefinition{
		Name: p.GetName(),
		Input: []*VariableDefinition{
			{
				Name: "left",
				Type: "string",
			},
			{
				Name: "right",
				Type: "string",
			},
		},
		Output: []*VariableDefinition{
			{
				Name: "return",
				Type: "string",
			},
		},
	}
}

type StringReplaceFunctionEntity struct {
}

func NewStringReplaceFunctionEntity() *StringReplaceFunctionEntity {
	return &StringReplaceFunctionEntity{}
}

func (p *StringReplaceFunctionEntity) GetName() string {
	return "StringReplace"
}
func (p *StringReplaceFunctionEntity) ToFunction(m map[string]interface{}) (*Function, error) {
	var option BaseFunctionOption
	err := mapstructure.Decode(m, &option)
	if err != nil {
		return nil, err
	}
	err = option.CheckValidate(p)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	function := &Function{
		Name: p.GetName(),
		OnRun: func(f *Function, runtime *Runtime) error {
			source, err := ReadVariable(f, "source", "string", &option)
			if err != nil {
				return err
			}
			search, err := ReadVariable(f, "search", "string", &option)
			if err != nil {
				return err
			}
			replace, err := ReadVariable(f, "replace", "string", &option)
			if err != nil {
				return err
			}
			result := strings.ReplaceAll(source.Value.(string), search.Value.(string), replace.Value.(string))
			runtime.ActiveSignalChan <- &ActiveSignal{Id: f.Id, Output: []*Variable{
				{
					Name:  "return",
					Value: result,
					Type:  "string",
				},
			}}
			return nil
		},
		Outputs: []*Variable{},
		OnActive: func(active *ActiveSignal, function *Function, runtime *Runtime) {
			DefaultOnActive(p.GetDefinition(), option, active, function)
		},
	}
	return function, nil

}

func (p *StringReplaceFunctionEntity) GetDefinition() *FunctionDefinition {
	return &FunctionDefinition{
		Name: p.GetName(),
		Input: []*VariableDefinition{
			{
				Name: "source",
				Type: "string",
			},
			{
				Name: "search",
				Type: "string",
			},
			{
				Name: "replace",
				Type: "string",
			},
		},
		Output: []*VariableDefinition{
			{
				Name: "return",
				Type: "string",
			},
		},
	}
}
