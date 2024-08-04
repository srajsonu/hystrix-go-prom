package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
	"time"
)

type Config struct {
	Timeout                int `env:"TIMEOUT_IN_MS"`
	MaxConcurrentRequests  int `env:"MAX_CONCURRENT_REQUESTS"`
	RequestVolumeThreshold int `env:"REQUEST_VOLUME_THRESHOLD"`
	SleepWindow            int `env:"SLEEP_WINDOW"`
	ErrorPercentThreshold  int `env:"ERROR_PERCENT_THRESHOLD"`
}

func setupHystrixConfig(command string, config Config) {
	hystrix.ConfigureCommand(command, hystrix.CommandConfig{
		Timeout:                int(time.Duration(config.Timeout) * time.Millisecond),
		MaxConcurrentRequests:  config.MaxConcurrentRequests,
		RequestVolumeThreshold: config.RequestVolumeThreshold,
		SleepWindow:            config.SleepWindow,
		ErrorPercentThreshold:  config.ErrorPercentThreshold,
	})
}
