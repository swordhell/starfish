package main

import (
	"github.com/sirupsen/logrus"
	"starfish/starfish/cmd"
)

func main() {
	logrus.Infoln("Hello, World!")
	cmd.Execute()
}
