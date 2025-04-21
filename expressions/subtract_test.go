package expressions

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSubtractCorrect(t *testing.T) {
	subtract := Subtract(2, 1)
	result := 1.0

	if subtract != result {
		t.Error("Resultado esperado: ", result, " Valor recebido: ", subtract)
	}
}

func TestShouldntSubtractWithInf(t *testing.T) {
	result := Subtract(math.Inf(1), 1)
	if !math.IsInf(result, 1) {
		t.Errorf("Esperado +Inf, mas obtido: %v", result)
	}

	assert.True(t, math.IsInf(result, 1), "Esperado +Inf mas obtido: %v", result)
}
