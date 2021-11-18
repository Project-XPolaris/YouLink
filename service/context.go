package service

import "context"

var DefaultServiceContext = ServiceContext{}

type ServiceContext struct {
	DefaultLauncher       *ProgramLauncher
	DefaultRuntime        *Runtime
	DefaultFunctionHub    *FunctionEntityHub
	DefaultProcessManager *ProcessManager
}

func (s *ServiceContext) Init(context context.Context) {
	s.DefaultLauncher = NewProgramLauncher()
	s.DefaultRuntime = NewRuntime()
	s.DefaultFunctionHub = NewFunctionHub()
	s.DefaultProcessManager = NewProcessManager()
	go s.DefaultRuntime.Run(context)
	go s.DefaultLauncher.Run(context, s)
	RegisterDefaultFunction(s.DefaultFunctionHub)
}
