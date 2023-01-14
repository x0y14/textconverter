package textconverter

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"strconv"
	"testing"
)

func TestI2S(t *testing.T) {
	for i := 0; i < 1000; i++ {
		random := rand.Intn(10000)
		assert.Equal(t, strconv.Itoa(random), i2s(random))
	}
}

func TestBeki(t *testing.T) {
	for i := 0; i < 1000; i++ {
		x := rand.Intn(10)
		n := rand.Intn(10)
		expect := math.Pow(float64(x), float64(n))
		assert.Equal(t, int(expect), beki(x, n))
	}
}

func TestS2I(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := rand.Intn(10000)
		assert.Equal(t, n, s2i(strconv.Itoa(n)))
	}
}
