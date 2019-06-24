package bus

import (
	"github.com/sirupsen/logrus"
)


type Bus interface {
	PublishMessage(string, *logrus.Entry) error
}