package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type balanceMock struct {
	mock.Mock
}

func (p balanceMock) getBalance(accNo string, sortCode string) (int, int) {
	args := p.Called(accNo, sortCode)
	return args.Int(0), args.Int(1)
}

func TestQueryBalance(t *testing.T) {
	theMock := balanceMock{}
	theMock.On("getBalance", "0000", "01-01-01").Return(1000, 2200)

	bal, avail := queryBalance(theMock, "0000", "01-01-01")
	assert.Equal(t, 1000, bal)
	assert.Equal(t, 2200, avail)
}
