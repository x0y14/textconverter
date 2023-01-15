package textconverter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
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

func TestBeki10Minus(t *testing.T) {
	for i := 0; i < 100; i++ {
		x := rand.Intn(6)
		assert.Equal(t, math.Pow(float64(10), float64(-x)), beki10Minus(-x))
	}
}

func TestS2I(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := rand.Intn(10000)
		assert.Equal(t, n, s2i(strconv.Itoa(n)))
	}
}

func TestI2F(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := rand.Intn(10000)
		assert.Equal(t, float64(n), i2f(n))
	}
}

//	func TestF2S(t *testing.T) {
//		f2s(1121.320099)
//	}
func random(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}

func TestS2F(t *testing.T) {
	//assert.Equal(t, float64(2), s2f("2"))
	assert.Equal(t, 1.2, s2f("1.2"))
	assert.Equal(t, 12.345, s2f("12.345"))

	for i := 0; i < 100; i++ {
		r := float64(rand.Intn(100)) + random(1.0, 10.0)
		rs := fmt.Sprintf("%f", r)
		f, err := strconv.ParseFloat(rs, 64)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, f, s2f(rs))
	}
}
