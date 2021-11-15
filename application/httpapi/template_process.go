package httpapi

import "github.com/projectxpolaris/youlink/service"

type ProcessProgram struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type BaseProcessTemplate struct {
	Id       string            `json:"id"`
	Programs []*ProcessProgram `json:"programs"`
}

func SerializerBaseProcessTemplate(process *service.Process) *BaseProcessTemplate {
	template := &BaseProcessTemplate{
		Id:       process.Id,
		Programs: []*ProcessProgram{},
	}
	for _, program := range process.Programs {
		programTemplate := &ProcessProgram{
			Id:   program.Id,
			Name: program.Name,
		}
		template.Programs = append(template.Programs, programTemplate)
	}
	return template
}
