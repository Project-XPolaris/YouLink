package httpapi

import "github.com/projectxpolaris/youlink/service"

type DefinitionVariableSerializerTemplate struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type FunctionSerializerTemplate struct {
	Name    string                                  `json:"name"`
	Inputs  []*DefinitionVariableSerializerTemplate `json:"inputs"`
	Outputs []*DefinitionVariableSerializerTemplate `json:"outputs"`
}

func SerializerFunctionDefinitionList(functions []service.FunctionEntity) []*FunctionSerializerTemplate {
	data := make([]*FunctionSerializerTemplate, 0)
	for _, function := range functions {
		definition := function.GetDefinition()
		template := &FunctionSerializerTemplate{
			Name:    definition.Name,
			Inputs:  []*DefinitionVariableSerializerTemplate{},
			Outputs: []*DefinitionVariableSerializerTemplate{},
		}
		for _, definition := range definition.Input {
			template.Inputs = append(template.Inputs, &DefinitionVariableSerializerTemplate{
				Name: definition.Name,
				Type: definition.Type,
			})
		}
		for _, definition := range definition.Output {
			template.Outputs = append(template.Outputs, &DefinitionVariableSerializerTemplate{
				Name: definition.Name,
				Type: definition.Type,
			})
		}
		data = append(data, template)
	}
	return data
}
