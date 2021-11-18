package service

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type PlusFunctionEntity struct {
}

func NewPlusFunctionEntity() *PlusFunctionEntity {
	return &PlusFunctionEntity{}
}

func (p *PlusFunctionEntity) GetName() string {
	return "Plus"
}
func (p *PlusFunctionEntity) ToFunction(m map[string]interface{}) (*Function, error) {
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
			left, err := ReadVariable(f, "left", "number", &option)
			if err != nil {
				return err
			}
			right, err := ReadVariable(f, "right", "number", &option)
			if err != nil {
				return err
			}
			result := left.Value.(float64) + right.Value.(float64)
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
		OnActive: func(active *ActiveSignal, function *Function, runtime *Runtime) {
			DefaultOnActive(p.GetDefinition(), option, active, function)
		},
	}
	return function, nil

}

func (p *PlusFunctionEntity) GetDefinition() *FunctionDefinition {
	return &FunctionDefinition{
		Name: p.GetName(),
		Input: []*VariableDefinition{
			{
				Name: "left",
				Type: "number",
			},
			{
				Name: "right",
				Type: "number",
			},
		},
		Output: []*VariableDefinition{
			{
				Name: "return",
				Type: "number",
			},
		},
	}
}

type SubtractFunctionEntity struct {
}

func NewSubtractFunctionEntity() *SubtractFunctionEntity {
	return &SubtractFunctionEntity{}
}

func (p *SubtractFunctionEntity) GetName() string {
	return "Sub"
}
func (p *SubtractFunctionEntity) ToFunction(m map[string]interface{}) (*Function, error) {
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
			left, err := ReadVariable(f, "left", "number", &option)
			if err != nil {
				return err
			}
			right, err := ReadVariable(f, "right", "number", &option)
			if err != nil {
				return err
			}
			result := left.Value.(float64) - right.Value.(float64)
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
		OnActive: func(active *ActiveSignal, function *Function, runtime *Runtime) {
			DefaultOnActive(p.GetDefinition(), option, active, function)
		},
	}
	return function, nil

}

func (p *SubtractFunctionEntity) GetDefinition() *FunctionDefinition {
	return &FunctionDefinition{
		Name: p.GetName(),
		Input: []*VariableDefinition{
			{
				Name: "left",
				Type: "number",
			},
			{
				Name: "right",
				Type: "number",
			},
		},
		Output: []*VariableDefinition{
			{
				Name: "return",
				Type: "number",
			},
		},
	}
}

type MultiFunctionEntity struct {
}

func NewMultiFunctionEntity() *MultiFunctionEntity {
	return &MultiFunctionEntity{}
}

func (p *MultiFunctionEntity) GetName() string {
	return "Mul"
}
func (p *MultiFunctionEntity) ToFunction(m map[string]interface{}) (*Function, error) {
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
			left, err := ReadVariable(f, "left", "number", &option)
			if err != nil {
				return err
			}
			right, err := ReadVariable(f, "right", "number", &option)
			if err != nil {
				return err
			}
			result := left.Value.(float64) * right.Value.(float64)
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
		OnActive: func(active *ActiveSignal, function *Function, runtime *Runtime) {
			DefaultOnActive(p.GetDefinition(), option, active, function)
		},
	}
	return function, nil

}

func (p *MultiFunctionEntity) GetDefinition() *FunctionDefinition {
	return &FunctionDefinition{
		Name: p.GetName(),
		Input: []*VariableDefinition{
			{
				Name: "left",
				Type: "number",
			},
			{
				Name: "right",
				Type: "number",
			},
		},
		Output: []*VariableDefinition{
			{
				Name: "return",
				Type: "number",
			},
		},
	}
}

type DivideFunctionEntity struct {
}

func NewDivideFunctionEntity() *DivideFunctionEntity {
	return &DivideFunctionEntity{}
}

func (p *DivideFunctionEntity) GetName() string {
	return "Div"
}
func (p *DivideFunctionEntity) ToFunction(m map[string]interface{}) (*Function, error) {
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
			left, err := ReadVariable(f, "left", "number", &option)
			if err != nil {
				return err
			}
			right, err := ReadVariable(f, "right", "number", &option)
			if err != nil {
				return err
			}
			result := left.Value.(float64) / right.Value.(float64)
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
		OnActive: func(active *ActiveSignal, function *Function, runtime *Runtime) {
			DefaultOnActive(p.GetDefinition(), option, active, function)
		},
	}
	return function, nil

}

func (p *DivideFunctionEntity) GetDefinition() *FunctionDefinition {
	return &FunctionDefinition{
		Name: p.GetName(),
		Input: []*VariableDefinition{
			{
				Name: "left",
				Type: "number",
			},
			{
				Name: "right",
				Type: "number",
			},
		},
		Output: []*VariableDefinition{
			{
				Name: "return",
				Type: "number",
			},
		},
	}
}
func ReadVariable(function *Function, name string, varType string, option *BaseFunctionOption) (*Variable, error) {
	inputVariable := option.GetInputByName(name)
	if inputVariable == nil {
		return nil, errors.New(fmt.Sprintf("input [%s] - [%s] not found", name, varType))
	}
	if len(inputVariable.Ref) > 0 {
		for _, variable := range function.Context {
			if variable.Name == inputVariable.Ref {
				return variable, nil
			}
		}
		return nil, errors.New(fmt.Sprintf("reference variable [%s] not found in context", name))
	}
	return &Variable{
		Name:  inputVariable.Name,
		Value: inputVariable.Value,
		Type:  varType,
	}, nil
}
