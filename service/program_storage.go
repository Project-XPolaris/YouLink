package service

import (
	"github.com/projectxpolaris/youlink/database"
)

func SaveProgram(name string, body []map[string]interface{}) (*database.Program, error) {
	program := &database.Program{
		Name: name,
	}
	err := program.SetProgramBody(body)
	if err != nil {
		return nil, err
	}
	err = database.DefaultDatabasePlugin.DB.Save(program).Error
	if err != nil {
		return nil, err
	}
	return program, nil
}
func CreateProcessInstance(id uint) (*Process, error) {
	var storeProgram database.Program
	err := database.DefaultDatabasePlugin.DB.Find(&storeProgram, id).Error
	if err != nil {
		return nil, err
	}
	process := DefaultProcessManager.AllocateProcess()
	program, err := CreateProgram(&storeProgram)
	if err != nil {
		return nil, err
	}
	err = process.AddProgram(program)
	if err != nil {
		return nil, err
	}
	return process, err
}
func CreateProgram(storeProgram *database.Program) (*Program, error) {
	body, err := storeProgram.GetProgramBody()
	if err != nil {
		return nil, err
	}
	program, err := Parse(body)
	if err != nil {
		return nil, err
	}
	program.Name = storeProgram.Name
	return program, nil
}
