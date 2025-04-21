package expressions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldDivideCorrect(t *testing.T) {
	divide, err := Divide(4, 2)

	result := 2.0
	if err == nil && result != divide {
		t.Error("Resultado esperado: ", result, " valor recebido: ", divide)
	}
}

func TestShouldntDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)

	assert.Error(t, err)

	assert.EqualError(t, err, "divisão por zero não permitida")
}
