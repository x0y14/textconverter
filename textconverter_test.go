package textconverter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestI2S(t *testing.T) {
	assert.Equal(t, "10", i2s(10))
}

func TestS2I(t *testing.T) {
	assert.Equal(t, 123, s2i("123"))
}

func TestBeki(t *testing.T) {
	assert.Equal(t, 32, beki(2, 5))
	assert.Equal(t, 1, beki(2, 0))
	assert.Equal(t, 2, beki(2, 1))
}
