package main

import (
	"github.com/allentom/harukap"
	"github.com/allentom/harukap/cli"
	"github.com/projectxpolaris/youlink/application/httpapi"
	"github.com/projectxpolaris/youlink/config"
	"github.com/projectxpolaris/youlink/database"
	"github.com/projectxpolaris/youlink/youlog"
	"github.com/projectxpolaris/youlink/youplus"
	"github.com/sirupsen/logrus"
)

func main() {
	err := config.InitConfigProvider()
	if err != nil {
		logrus.Fatal(err)
	}
	err = youlog.DefaultYouLogPlugin.OnInit(config.DefaultConfigProvider)
	if err != nil {
		logrus.Fatal(err)
	}

	appEngine := harukap.NewHarukaAppEngine()
	appEngine.ConfigProvider = config.DefaultConfigProvider
	appEngine.LoggerPlugin = youlog.DefaultYouLogPlugin
	appEngine.UsePlugin(&youplus.DefaultYouPlusPlugin)
	appEngine.UsePlugin(&database.DefaultDatabasePlugin)
	appEngine.UsePlugin(&ApplicationInitPlugin{})
	appEngine.HttpService = httpapi.GetEngine()
	appWrap, err := cli.NewWrapper(appEngine)
	if err != nil {
		logrus.Fatal(err)
	}
	appWrap.RunApp()
}
