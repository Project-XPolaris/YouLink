package service

import (
	"sync"
)

func RegisterDefaultFunction(hub *FunctionEntityHub) {
	hub.RegisterFunctions(
		NewUpdatePathDirFunction(),
		NewPlusFunctionEntity(),
		NewSubtractFunctionEntity(),
		NewMultiFunctionEntity(),
		NewDivideFunctionEntity(),
		NewConcatFunctionEntity(),
	)
}

type VariableDefinition struct {
	Type string
	Name string
}
type FunctionDefinition struct {
	Name   string
	Input  []*VariableDefinition
	Output []*VariableDefinition
}

func (f *FunctionDefinition) GetOutputDefinitionByName(name string) *VariableDefinition {
	for _, definition := range f.Output {
		if definition.Name == name {
			return definition
		}
	}
	return nil
}

type FunctionEntity interface {
	GetName() string
	ToFunction(m map[string]interface{}) (*Function, error)
	GetDefinition() *FunctionDefinition
}
type FunctionEntityHub struct {
	sync.Mutex
	Functions []FunctionEntity
}

func NewFunctionHub() *FunctionEntityHub {
	return &FunctionEntityHub{Functions: []FunctionEntity{}}
}

func (h *FunctionEntityHub) RegisterFunctions(functions ...FunctionEntity) {
	h.Lock()
	defer h.Unlock()
	h.Functions = append(h.Functions, functions...)
}

func (h *FunctionEntityHub) GetEntityByName(name string) FunctionEntity {
	for _, function := range h.Functions {
		if function.GetName() == name {
			return function
		}
	}
	return nil
}

type FunctionInput struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Ref   string      `json:"ref"`
}
type FunctionOutput struct {
	Name   string `json:"name"`
	Assign string `json:"assign"`
}

func (f *UpdatePathDirFunction) GetName() string {
	return f.Definition.Name
}
