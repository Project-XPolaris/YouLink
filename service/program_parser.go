package service

import (
	"errors"
	"fmt"
	"github.com/projectxpolaris/youlink/utils"
)

type ProgramBody struct {
	Body []map[string]interface{} `json:"body"`
}

func Parse(body []map[string]interface{}, serviceContext ServiceContext) (*Program, error) {
	// for body
	program := NewProgram()
	// find trigger
	entryBlock, err := NewRootBlock(serviceContext)
	if err != nil {
		return nil, err
	}
	err = walk(entryBlock, body, serviceContext)
	if err != nil {
		return nil, err
	}
	program.Runners = append(program.Runners, entryBlock)
	return program, nil
}
func walk(parenBlock *Block, body []map[string]interface{}, serviceContext ServiceContext) error {
	for _, value := range body {
		if blockType, exist := value["type"]; exist {
			if blockType == "function" {
				function, err := GenerateFunction(value, serviceContext)
				if err != nil {
					return err
				}
				parenBlock.Runners = append(parenBlock.Runners, function)
			}
			if blockType == "block" {
				blockBody := utils.MapConvert(value["body"].([]interface{}))
				block, err := NewBlock(value, serviceContext)
				if err != nil {
					return err
				}
				err = walk(block, blockBody, serviceContext)
				if err != nil {
					return err
				}
				parenBlock.Runners = append(parenBlock.Runners, block)
			}
		}
	}
	return nil
}
func GenerateFunction(value map[string]interface{}, serviceContext ServiceContext) (*Function, error) {
	name := value["name"]
	functionEntity := serviceContext.DefaultFunctionHub.GetEntityByName(name.(string))
	if functionEntity == nil {
		return nil, errors.New(fmt.Sprintf("function [%s] not found", name))
	}
	function, err := functionEntity.ToFunction(value)
	if err != nil {
		return nil, err
	}
	return function, nil
}
