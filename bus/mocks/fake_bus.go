package mocks

import (
	"github.com/stretchr/testify/mock"
)

type FakeBus struct {
	mock.Mock
}

func (b *FakeBus) SendMessage(msg []byte) error {
	args := b.Called(msg)
	return args.Error(0)
}

func (b *FakeBus) ConsumeMessages() error {
	args := b.Called()
	return args.Error(0)
}
