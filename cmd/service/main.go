package main

import (
	"github.com/none-da/swappable-logging/pkg/service"
	flag "github.com/spf13/pflag"
)

func main() {
	var requiredLogger string
	flag.StringVarP(&requiredLogger, "logger", "l", "logrus", "Required logger. logrus/zap.")
	var printHelp bool
	flag.BoolVarP(&printHelp, "help", "h", false, "Prints this help content.")
	flag.Parse()
	if printHelp {
		flag.Usage()
		return
	}

	svc := service.CreateService(requiredLogger)
	service.DoSomething(&svc, "Serve the traffic")
}
