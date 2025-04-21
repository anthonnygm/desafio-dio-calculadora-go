package expressions

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldMultiplyCorrect(t *testing.T) {
	multiply := Multiply(2, 3)

	result := 6.0

	assert.Equal(t, result, multiply)
}

func TestShouldntMultiplyWithInf(t *testing.T) {
	result := Multiply(math.Inf(1), 0)

	if !math.IsNaN(result) {
		t.Errorf("Multiplicação de zero por infinito deve resultar em NaN, mas obtido: %v", result)
	}

	assert.True(t, math.IsNaN(result), "Multiplicação de zero por infinito deve resultar em NaN, mas obtido: %v", result)
}
