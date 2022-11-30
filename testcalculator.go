package testcalculator

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		num1     int
		num2     int
		operator string
		solution int
	}{
		{25, 25, "+", 100},
		{100, 200, "+", 300},
		{10, 5, "+", 15},
		{25, 10, "-", 15},
		{750, 800, "-", -150},
		{35, 15, "-", 20},
		{10, 10, "*", 100},
		{300, 1, "*", 300},
		{12, 0, "*", 0},
		{45, 5, "/", 9},
		{500, 100, "/", 5},
		{12, 0, "/", 0},
		{12, 12, "^", 144},
		{0, 0, "^", 1},
		{0, 0, "^", 1},
	}

	for _, cases := range testCases {
		t.Run(fmt.Sprintf("%d%s%d", cases.num1, cases.operator, cases.num2), func(t *testing.T) {
			switch cases.operator {
			case "+":
				if result := calculator.Add(cases.num1, cases.num2); result != cases.solution {
					t.Errorf("result: %d - solution: %d", result, cases.solution)
				}
			case "-":
				if result := calculator.Subtract(cases.num1, cases.num2); result != cases.solution {
					t.Errorf("result: %d - solution: %d", result, cases.solution)
				}
			case "*":
				if result := calculator.Multiply(cases.num1, cases.num2); result != cases.solution {
					t.Errorf("result: %d - solution: %d", result, cases.solution)
				}
			case "/":
				if result, err := calculator.Divide(cases.num1, cases.num2); err != nil {
					if cases.num2 != 0 {
						t.Errorf("%d is not a number", cases.num2)
					}
				} else if result != cases.solution {
					t.Errorf("result: %d - solution: %d", result, cases.solution)
				}
			case "^":
				if result := calculator.Pow(float64(cases.num1), float64(cases.num2)); result != float64(cases.solution) {
					t.Errorf("result: %f - solution: %d", result, cases.solution)
				}
			default:
				t.Errorf("invalid: %s", cases.operator)
			}
		})
	}
}
