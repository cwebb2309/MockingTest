package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type balanceMock struct {
	mock.Mock
}

func (p balanceMock) getBalance(accNo string) int {
	args := p.Called(accNo)
	return args.Int(0)
}

func TestMockMethodWithArgs(t *testing.T) {
	theMock := balanceMock{}
	theMock.On("getBalance", "0000").Return(1000, 0)
	res := queryBalance(theMock, "0000")
	assert.Equal(t, 1000, res)
	//theMock.AssertExpectations(t)
}
