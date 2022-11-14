package testcalculator

import (
	"fmt"
	"testing"
)

func main()
func TestCalculator(t *testing.T) {
	testCases := []struct {
		num1     int
		num2     int
		operator string
		solution int
	}{
		{15, 27, "+", 42},
		{957, 767, "+", 1724},
		{7, 4, "+", 11},
		{30, 15, "-", 15},
		{500, 700, "-", -200},
		{15, 10, "-", 5},
		{25, 25, "*", 625},
		{200, 1, "*", 200},
		{25, 0, "*", 0},
		{14, 7, "/", 2},
		{1000, 25, "/", 40},
		{10, 0, "/", 0},
	}

	for _, cases := range testCases {
		t.Run(fmt.Sprintf("%d%s%d", cases.num1, cases.operator, cases.num2), func(t *testing.T) {
			switch cases.operator {
			case "+":
				if result := TestCalculator.Add(cases.num1, cases.num2); result != cases.solution {
					t.Errorf("result: %d - solution: %d", result, cases.solution)
				}
			case "-":
				if result := TestCalculator.Subtract(cases.num1, cases.num2); result != cases.solution {
					t.Errorf("result: %d - solution: %d", result, cases.solution)
				}
			case "*":
				if result := TestCalculator.Multiply(cases.num1, cases.num2); result != cases.solution {
					t.Errorf("result: %d - solution: %d", result, cases.solution)
				}
			case "/":
				if result, err := TestCalculator.Divide(cases.num1, cases.num2); err != nil {
					if cases.num2 != 0 {
						t.Errorf("%d is not a number", cases.num2)
					}
				} else if result != cases.solution {
					t.Errorf("result: %d - solution: %d", result, cases.solution)
				}
			default:
				t.Errorf("invalid: %s", cases.operator)
			}
		})
	}
}
