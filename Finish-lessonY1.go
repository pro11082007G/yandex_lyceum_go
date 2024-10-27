package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func applyOp(a, b float64, op rune) float64 {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		if b != 0 {
			return a / b
		}
		panic("division by zero")
	}
	return 0
}

func Calc(expression string) (float64, error) {
	var values []float64
	var ops []rune


	expression = strings.ReplaceAll(expression, " ", "")

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])


		if unicode.IsDigit(ch) {
			var numStr string
			for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				numStr += string(expression[i])
				i++
			}
			i-- 
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", numStr)
			}
			values = append(values, num)
		} else if ch == '(' {
			ops = append(ops, ch)
		} else if ch == ')' {

			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if len(values) < 2 {
					return 0, errors.New("invalid expression: insufficient values for operation")
				}
				val2 := values[len(values)-1]
				values = values[:len(values)-1]
				val1 := values[len(values)-1]
				values = values[:len(values)-1]
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]

				values = append(values, applyOp(val1, val2, op))
			}
			if len(ops) == 0 {
				return 0, errors.New("mismatched parentheses: too many closing parentheses")
			}
			ops = ops[:len(ops)-1] // Удаляем '('
		} else if precedence(ch) > 0 {

			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(ch) {
				if len(values) < 2 {
					return 0, errors.New("invalid expression: insufficient values for operation")
				}
				val2 := values[len(values)-1]
				values = values[:len(values)-1]
				val1 := values[len(values)-1]
				values = values[:len(values)-1]
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]

				values = append(values, applyOp(val1, val2, op))
			}
			ops = append(ops, ch)
		} else if !unicode.IsSpace(ch) {
			return 0, fmt.Errorf("invalid character: %c", ch)
		}
	}


	for len(ops) > 0 {
		if len(values) < 2 {
			return 0, errors.New("invalid expression: insufficient values for operation")
		}
		val2 := values[len(values)-1]
		values = values[:len(values)-1]
		val1 := values[len(values)-1]
		values = values[:len(values)-1]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		values = append(values, applyOp(val1, val2, op))
	}

	if len(values) == 1 {
		return values[0], nil
	}
	return 0, errors.New("invalid expression: unmatched operators and values")
}
