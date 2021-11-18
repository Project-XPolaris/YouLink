package localtest

//import (
//	"context"
//	"encoding/json"
//	"github.com/projectxpolaris/youlink/application/httpapi"
//	"github.com/projectxpolaris/youlink/service"
//	"io/ioutil"
//	"testing"
//)
//
//func TestRuntime(t *testing.T) {
//	go service.DefaultRuntime.Run(context.Background())
//	go service.DefaultLauncher.Run(context.Background())
//	program := service.NewProgram()
//	program.Context = append(program.Context, &service.Variable{
//		Name:  "left",
//		Value: 1,
//		Type:  "number",
//	}, &service.Variable{
//		Name:  "right",
//		Value: 2,
//		Type:  "number",
//	})
//	plusFunction := service.NewPlusFunction()
//	program.Runners = append(program.Runners, plusFunction)
//	service.DefaultLauncher.Queue <- program
//	<-program.OnDone
//	t.Log(program.Context)
//}
//func TestYouFileMove(t *testing.T) {
//	go func() {
//		e := httpapi.GetEngine()
//		e.RunAndListen(":4700")
//	}()
//	go service.DefaultRuntime.Run(context.Background())
//	go service.DefaultLauncher.Run(context.Background())
//	service.DefaultYouFileClient.Init("http://localhost:8300")
//	program := service.NewProgram()
//	program.Context = append(program.Context, &service.Variable{
//		Name:  "source",
//		Value: "C:\\Users\\aren\\Desktop\\video_library\\video1.mp4",
//		Type:  "string",
//	}, &service.Variable{
//		Name:  "target",
//		Value: "C:\\Users\\aren\\Desktop\\download\\video1.mp4",
//		Type:  "string",
//	})
//	plusFunction := service.NewYouFileMoveFunction()
//	program.Runners = append(program.Runners, plusFunction)
//	service.DefaultLauncher.Queue <- program
//	<-program.OnDone
//	if program.Error != nil {
//		t.Error(program.Error)
//	}
//	t.Log(program.Id)
//}
//
//func TestParse(t *testing.T) {
//	go func() {
//		e := httpapi.GetEngine()
//		e.RunAndListen(":4700")
//	}()
//	go service.DefaultRuntime.Run(context.Background())
//	go service.DefaultLauncher.Run(context.Background())
//	service.RegisterDefaultFunction(service.DefaultFunctionHub)
//
//	planRegData, _ := ioutil.ReadFile("../example/reg_func.json")
//	var regData httpapi.RegisterFunctionsRequestBody
//	err := json.Unmarshal(planRegData, &regData)
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = service.ParseFunctionTemplate(regData.Func)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	plan, _ := ioutil.ReadFile("../example/make.json")
//	var data service.ProgramBody
//	err = json.Unmarshal(plan, &data)
//	if err != nil {
//		t.Fatal(err)
//	}
//	program, err := service.Parse(data.Body)
//	if err != nil {
//		t.Fatal(err)
//	}
//	program.Context = append(program.Context, &service.Variable{
//		Name:  "youdownloadSavePath",
//		Value: "C:\\Users\\aren\\Desktop\\video_library\\video1.mp4",
//		Type:  "string",
//	})
//	service.DefaultLauncher.Queue <- program
//	<-program.OnDone
//
//	t.Log(program)
//}
