package service

import (
	"errors"
	"fmt"
)

type ProgramBody struct {
	Body []map[string]interface{} `json:"body"`
}

func Parse(body []map[string]interface{}) (*Program, error) {
	// for body
	program := NewProgram()
	// find trigger
	for _, value := range body {
		if blockType, exist := value["type"]; exist {
			if blockType == "function" {
				name := value["name"]
				functionEntity := DefaultFunctionHub.GetEntityByName(name.(string))
				if functionEntity == nil {
					return nil, errors.New(fmt.Sprintf("function [%s] not found", name))
				}
				function, err := functionEntity.ToFunction(value)
				if err != nil {
					return nil, err
				}
				program.Functions = append(program.Functions, function)
			}
		}
	}
	return program, nil
}
