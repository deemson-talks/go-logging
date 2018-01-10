package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	sublogger := logrus.WithFields(logrus.Fields{"param1": "value1"})
	sublogger.WithField("param2", "value2").Info("hello world")
}
