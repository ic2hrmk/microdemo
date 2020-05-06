package main

import (
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"

	"github.com/ic2hrmk/microdemo/microdemo-web/handler"
)

const (
	DefaultServiceName = "dev.ic2hrmk.api.microdemo"
)

func main() {
	var (
		showVersionOnlyConfig bool
		handlerPathConfig     string
	)

	var (
		handlerPathFlag = cli.StringFlag{
			Name:        "handler",
			Value:       "/api/call",
			Usage:       "Path to receive HTTP requests",
			EnvVar:      "MICRODEMO_WEB_HANDLER",
			Destination: &handlerPathConfig,
		}
		showVersionFlag = cli.BoolFlag{
			Name:        "version",
			Usage:       "Shows a version of application and exists",
			Hidden:      false,
			Destination: &showVersionOnlyConfig,
		}
	)

	service := web.NewService(
		web.Name(DefaultServiceName),
		web.Version("latest"),
		web.Flags(
			handlerPathFlag,
			showVersionFlag,
		),
	)

	if err := service.Init()
		err != nil {
		log.Fatal(err)
	}

	if showVersionOnlyConfig == true {
		log.Infof("microdemo-web, version: %s", service.Options().Version)
		os.Exit(0)
	}

	service.HandleFunc(handlerPathConfig, handler.MicrodemoCall(service))

	log.Infof("Hello, I am %s [version=%s], ready to accept http requests at [%s]",
		service.Options().Name,
		service.Options().Version,
		handlerPathConfig,
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
