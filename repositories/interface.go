package repositories

import (
	"github.com/sirupsen/logrus"
)

type Repository interface {
	PublishMessage(string, string, *logrus.Entry) error
}
