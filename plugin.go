package main

import (
	"context"
	"github.com/allentom/harukap"
	"github.com/projectxpolaris/youlink/service"
)

type ApplicationInitPlugin struct {
}

func (p *ApplicationInitPlugin) OnInit(e *harukap.HarukaAppEngine) error {
	service.DefaultServiceContext.Init(context.Background())
	return nil
}
