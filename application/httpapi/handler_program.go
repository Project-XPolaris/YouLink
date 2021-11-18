package httpapi

import (
	"errors"
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/youlink/service"
	"net/http"
)

type CreateProgramRequestBody struct {
	Name string                   `json:"name"`
	Body []map[string]interface{} `json:"body"`
}

var createProgramHandler haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody CreateProgramRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	program, err := service.SaveProgram(requestBody.Name, requestBody.Body)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	template, _ := NewProgramSerializerTemplate(program)
	context.JSON(haruka.JSON{
		"success": true,
		"data":    template,
	})
}

type NewProgramInstanceRequestBody struct {
	Id uint `json:"id"`
}

var newInstanceProgram haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody NewProgramInstanceRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	process, err := service.CreateProcessInstance(requestBody.Id, service.DefaultServiceContext)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	data := SerializerBaseProcessTemplate(process)
	context.JSON(haruka.JSON{
		"success": true,
		"data":    data,
	})
}

type RunProcessRequestBody struct {
	Id     string              `json:"id"`
	Inputs []*service.Variable `json:"inputs"`
}

var runProcessHandler haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody RunProcessRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	process := service.DefaultServiceContext.DefaultProcessManager.GetProcessById(requestBody.Id)
	if process == nil {
		AbortError(context, errors.New("process not found"), http.StatusNotFound)
		return
	}
	err = process.SetInput(requestBody.Inputs)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	process.Run(service.DefaultServiceContext)
	data := SerializerBaseProcessTemplate(process)
	context.JSON(haruka.JSON{
		"success": true,
		"data":    data,
	})
}
