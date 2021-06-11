package controller

import (
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

type MockReader struct {
	Counter int64
}

func (mr *MockReader) Read(data []byte) (int, error) {
	for i := 0; i < len(data); i++ {
		data[i] = byte(atomic.AddInt64(&mr.Counter, 1))
	}

	return len(data), nil
}

func TestNewGenerator(t *testing.T) {
	length := 4

	gen := Generator{Reader: &MockReader{}}
	gen.Length = &length
	gen.Numbers = true

	str, err := gen.Generate()
	assert.NoError(t, err)
	assert.Equal(t, "1234", str)
}

func TestExtraSymbols(t *testing.T) {
	length := 4

	gen := Generator{Reader: &MockReader{}}
	gen.Length = &length
	gen.ExtraSymbols = "_=+"

	str, err := gen.Generate()
	assert.NoError(t, err)
	assert.Equal(t, "=+_=", str)
}
