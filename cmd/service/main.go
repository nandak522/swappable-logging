package main

import (
	"github.com/none-da/swappable-logging/pkg/service"
	flag "github.com/spf13/pflag"
)

func main() {
	var requiredLogLevel string
	flag.StringVarP(&requiredLogLevel, "log-level", "l", "info", "Required log level. debug/info/warn/error.")
	var printHelp bool
	flag.BoolVarP(&printHelp, "help", "h", false, "Prints this help content.")
	flag.Parse()
	if printHelp {
		flag.Usage()
		return
	}

	svc := service.CreateService()
	service.DoSomething(&svc, "Serve the traffic")
}
