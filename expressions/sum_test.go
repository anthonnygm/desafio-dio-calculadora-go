package expressions

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSumCorrect(t *testing.T) {
	sum := Sum(1, 2)

	result := 3.0
	if sum != result {
		t.Error("Resultado esperado: ", result, "valor recebido: ", sum)
	}
}

func TestShouldntSumWithNaN(t *testing.T) {
	result := Sum(math.NaN(), 1)
	if !math.IsNaN(result) {
		t.Errorf("Esperado NaN, mas obtido: %v", result)
	}

	assert.True(t, math.IsNaN(result), "Esperado NaN, mas obtido: %v", result)
}
