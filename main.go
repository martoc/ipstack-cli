package main

import (
	"github.com/martoc/ipstack-cli/cmd"
	"github.com/martoc/ipstack-cli/logger"
)

func main() {
	cmd.Execute()
	logger.Close()
}
