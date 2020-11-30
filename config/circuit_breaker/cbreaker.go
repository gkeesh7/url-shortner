package circuit_breaker

import (
	"github.com/afex/hystrix-go/hystrix"
	"url-shortner/common"
)

var hystrixCommands = []string{
	common.READ_FROM_DATABASE,
}

func init() {
	InitializeHystrix()
}

func InitializeHystrix() {
	for _, element := range hystrixCommands {
		hystrix.ConfigureCommand(element, hystrix.CommandConfig{
			Timeout:               100,
			MaxConcurrentRequests: 2000,
			SleepWindow:           5000,
			ErrorPercentThreshold: 25,
		})
	}
}
