package database

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	Name string `json:"name"`
	Body string `json:"body"`
}

func (p *Program) GetProgramBody() ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	err := json.Unmarshal([]byte(p.Body), &data)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (p *Program) SetProgramBody(data []map[string]interface{}) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}
	p.Body = string(raw)
	return nil
}
