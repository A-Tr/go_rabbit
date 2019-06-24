package bus

import (
	"github.com/sirupsen/logrus"
)


type Bus interface {
	PublishMessage(string, string, *logrus.Entry) error
}