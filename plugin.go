package main

import (
	"context"
	"github.com/allentom/harukap"
	"github.com/projectxpolaris/youlink/service"
)

type ApplicationInitPlugin struct {
}

func (p *ApplicationInitPlugin) OnInit(e *harukap.HarukaAppEngine) error {
	go service.DefaultLauncher.Run(context.Background())
	go service.DefaultRuntime.Run(context.Background())
	service.RegisterDefaultFunction(service.DefaultFunctionHub)
	return nil
}
