package httpapi

import (
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/youlink/service"
)

var functionListHandler haruka.RequestHandler = func(context *haruka.Context) {
	data := SerializerFunctionDefinitionList(service.DefaultServiceContext.DefaultFunctionHub.Functions)
	context.JSON(haruka.JSON{
		"success": true,
		"data":    data,
	})
}
