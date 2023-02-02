package calculation

import (
	"log"
	"strconv"
	"unicode"
)

// Вычесляем значние операции
func opertaor(operator string, operandRight, operandLeft string) string {
	var resultStr string
	resultInt := 0
	operandRightInt, _ := strconv.Atoi(operandRight)
	operandLeftInt, _ := strconv.Atoi(operandLeft)

	switch operator {
	case "+":
		resultInt = operandRightInt + operandLeftInt
	case "-":
		resultInt = operandRightInt - operandLeftInt
	case "*":
		resultInt = operandRightInt * operandLeftInt
	case "/":
		if operandLeftInt == 0 {
			log.Fatal("Деление на 0")
		} else {
			resultInt = operandRightInt / operandLeftInt
		}
	}
	resultStr = strconv.Itoa(resultInt)
	return resultStr
}

// Calculate Польская нотация
func Calculate(example []rune) string {
	var result string
	var currentElement string
	stack := make([]string, len(example))
	list := make([]string, len(example))
	exampleString := string(example)
	_ = exampleString
	status := map[string]int{
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
	}
	stackCount := 0
	listCount := 0
	for i := 0; i < len(example); i++ {
		if unicode.IsNumber(example[i]) {
			list[listCount] = string(example[i])
			listCount++
		} else {
			stack[stackCount] += string(example[i])
			stackCount++
		}
		if stackCount >= 2 && status[stack[stackCount-1]] <= status[stack[stackCount-2]] {
			for stackCount >= 2 && status[stack[stackCount-1]] <= status[stack[stackCount-2]] {
				currentElement = stack[stackCount-1]
				stack[stackCount-1] = ""
				stackCount--
				listCount--
				result = opertaor(stack[stackCount-1], list[listCount-1], list[listCount])
				list[listCount] = ""
				list[listCount-1] = result
				stack[stackCount-1] = currentElement
			}
		}
	}
	if stackCount == 1 && stack[0] != "" {
		result = opertaor(stack[stackCount-1], list[listCount-2], list[listCount-1])
		list[listCount-1] = ""
		list[listCount-2] = result
		return list[0]
	}
	return list[0]
}
