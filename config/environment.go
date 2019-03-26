package config

import (
	"os"
)

var (
	BusType = os.Getenv("BUS_TYPE")
)