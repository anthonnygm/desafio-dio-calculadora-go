package expressions

import "errors"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não permitida")
	}

	return a / b, nil
}
