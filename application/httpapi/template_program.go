package httpapi

import "github.com/projectxpolaris/youlink/database"

type ProgramSerializerTemplate struct {
	Id   uint                     `json:"id"`
	Name string                   `json:"name"`
	Body []map[string]interface{} `json:"body"`
}

func NewProgramSerializerTemplate(program *database.Program) (*ProgramSerializerTemplate, error) {
	var err error
	template := &ProgramSerializerTemplate{
		Id:   program.ID,
		Name: program.Name,
	}
	template.Body, err = program.GetProgramBody()
	if err != nil {
		return nil, err
	}
	return template, err
}
