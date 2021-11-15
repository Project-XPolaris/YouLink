package service

import (
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
)

type HttpFunction struct {
	Definition *FunctionDefinition
	Option     *HttpFunctionTemplateOption
}

func (f *HttpFunction) GetName() string {
	return f.Definition.Name
}

func (f *HttpFunction) ToFunction(m map[string]interface{}) (*Function, error) {
	var option BaseFunctionOption
	err := mapstructure.Decode(m, &option)
	if err != nil {
		return nil, err
	}
	function := &Function{
		Name: f.Definition.Name,
		OnRun: func(function *Function, runtime *Runtime) error {
			inputVars := make([]*Variable, 0)
			for _, definition := range f.Definition.Input {
				for _, input := range option.Inputs {
					if input.Name == definition.Name {
						if len(input.Ref) > 0 {
							inputVars = append(inputVars, &Variable{
								Name:  definition.Name,
								Value: function.getInputByName(input.Ref).Value,
								Type:  definition.Type,
							})
						} else {
							inputVars = append(inputVars, &Variable{
								Name:  definition.Name,
								Value: input.Value.(string),
								Type:  definition.Type,
							})
						}
					}
				}
			}
			_, err := resty.New().R().
				SetBody(map[string]interface{}{
					"callbackId": function.Id,
					"inputs":     inputVars,
				}).
				Post(f.Option.Url)
			if err == nil {
				return err
			}
			return nil
		},
		OnActive: func(active *ActiveSignal, function *Function, runtime *Runtime) {
			DefaultOnActive(f.Definition, option, active, function)
		},
	}
	return function, nil
}
