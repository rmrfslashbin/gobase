package cmd

import (
	"github.com/sirupsen/logrus"
)

const PROGRAM_NAME = "gobase"
const VERSION = "0.0.1"

var (
	log       *logrus.Logger
	configDir string
	//Version   string
)
