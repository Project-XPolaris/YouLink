package httpapi

import (
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/youlink/service"
	"net/http"
)

type RegisterFunctionsRequestBody struct {
	Func []*service.BaseFunctionTemplate `json:"func"`
}

var registerFunctionHandler haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody RegisterFunctionsRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusBadRequest)
		return
	}
	err = service.ParseFunctionTemplate(requestBody.Func, &service.DefaultServiceContext)
	if err != nil {
		AbortErrorWithStatus(err, context, http.StatusBadRequest)
		return
	}
	context.JSON(haruka.JSON{
		"success": true,
	})
}
