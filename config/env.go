package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	LogLevelEnvVar       = "LOG_LEVEL"
	PortEnvVar           = "PORT"
	BrokerAddressEnvVar  = "BROKER_ADDRESS"
	defaultLogLevel      = logrus.DebugLevel
	defaultPort          = 8080
	defaultBrokerAddress = "localhost"
)

func LogLevel() logrus.Level {
	var (
		level logrus.Level
		err   error
	)

	if level, err = logrus.ParseLevel(os.Getenv(LogLevelEnvVar)); err != nil {
		return defaultLogLevel
	}

	return level
}

func Port() int {
	var (
		rawPort string
		found   bool
		port    int
		err     error
	)

	if rawPort, found = os.LookupEnv(PortEnvVar); !found {
		return defaultPort
	}

	if port, err = strconv.Atoi(rawPort); err != nil {
		return defaultPort
	}

	return port
}

func BrokerAddress() string {
	var (
		brokerAddress string
		found         bool
	)

	if brokerAddress, found = os.LookupEnv(BrokerAddressEnvVar); !found {
		return defaultBrokerAddress
	}

	if len(brokerAddress) == 0 {
		return defaultBrokerAddress
	}

	return brokerAddress
}
