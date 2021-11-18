package httpapi

import (
	"errors"
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/youlink/service"
	"net/http"
)

type YouLinkCallback struct {
	Id     string              `json:"id"`
	Output []*service.Variable `json:"output"`
	Error  string              `json:"error"`
}

var outputCallbackHandler haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody YouLinkCallback
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	activeSignal := &service.ActiveSignal{
		Id:     requestBody.Id,
		Output: requestBody.Output,
	}
	if len(requestBody.Error) > 0 {
		activeSignal.Error = errors.New(requestBody.Error)
	}
	service.DefaultServiceContext.DefaultRuntime.ActiveSignalChan <- activeSignal
	context.JSON(haruka.JSON{
		"success": true,
	})
}
