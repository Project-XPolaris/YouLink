package httpapi

import (
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/youlink/youlog"
)

func AbortError(ctx *haruka.Context, err error, status int) {
	youlog.DefaultYouLogPlugin.Logger.Error(err.Error())
	ctx.JSONWithStatus(haruka.JSON{
		"success": false,
		"reason":  err.Error(),
	}, status)
}
