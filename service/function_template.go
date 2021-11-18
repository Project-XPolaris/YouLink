package service

import (
	"github.com/mitchellh/mapstructure"
)

type BaseFunctionTemplate struct {
	Name     string                 `json:"name"`
	Template string                 `json:"template"`
	Desc     string                 `json:"desc"`
	Inputs   []*TemplateVariable    `json:"inputs"`
	Outputs  []*TemplateVariable    `json:"outputs"`
	Options  map[string]interface{} `json:"options"`
}
type TemplateVariable struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Type string `json:"type"`
}

func (v *TemplateVariable) ToVariableDefinition() *VariableDefinition {
	return &VariableDefinition{
		Type: v.Type,
		Name: v.Name,
	}
}
func (t *BaseFunctionTemplate) GetOutputVariableDefinition() []*VariableDefinition {
	definitions := make([]*VariableDefinition, 0)
	for _, variable := range t.Outputs {
		definitions = append(definitions, variable.ToVariableDefinition())
	}
	return definitions
}
func (t *BaseFunctionTemplate) GetInputVariableDefinition() []*VariableDefinition {
	definitions := make([]*VariableDefinition, 0)
	for _, variable := range t.Inputs {
		definitions = append(definitions, variable.ToVariableDefinition())
	}
	return definitions
}

func ParseFunctionTemplate(funcs []*BaseFunctionTemplate, serviceContext *ServiceContext) error {
	for _, template := range funcs {
		if template.Template == "HTTPRequestCall" {
			httpTemplate := HttpFunctionTemplate{
				Base: template,
			}
			option := &HttpFunctionTemplateOption{}
			err := mapstructure.Decode(&template.Options, option)
			if err != nil {
				return err
			}
			httpTemplate.Option = option
			entity, err := httpTemplate.Generate()
			if err != nil {
				return err
			}
			serviceContext.DefaultFunctionHub.RegisterFunctions(entity)
		}
	}
	return nil
}

type HttpFunctionTemplate struct {
	Base   *BaseFunctionTemplate
	Option *HttpFunctionTemplateOption `json:"option"`
}
type HttpFunctionTemplateOption struct {
	Url string `json:"url"`
}

func (t *HttpFunctionTemplate) Generate() (FunctionEntity, error) {
	entity := &HttpFunction{
		Definition: &FunctionDefinition{
			Name:   t.Base.Name,
			Input:  t.Base.GetInputVariableDefinition(),
			Output: t.Base.GetOutputVariableDefinition(),
		},
		Option: t.Option,
	}
	return entity, nil
}
