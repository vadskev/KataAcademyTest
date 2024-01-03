package internal

import (
	"strconv"
)

type NumberType int

const (
	Roman NumberType = iota
	Arabic
	Notype
)

const (
	NEGATIVENUMBER   = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	DIFFERENTNUMBERS = "Вывод ошибки, так как используются одновременно разные системы счисления."
	NOOPERATOR       = "Cтрока не является математической операцией."
	FORMAT           = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	NOVALID          = "Один операнд не является числом."
	NOVALIDRANGE     = "Числа выходят из диапазона от 1 до 10."
)

func Calc(line string) (string, bool) {
	isValidFormula := ValidateFormula(line)
	if isValidFormula {
		line, operator := FindOperation(line)
		numFirstString, numTwoString := SplitByOperator(line, operator)

		numTypeFirst := WhatIsNumberType(numFirstString)
		numTypeTwo := WhatIsNumberType(numTwoString)

		if numTypeFirst == Roman && numTypeTwo == Roman {
			result, isSuccess := ConvertAndCalculate(numFirstString, numTwoString, operator, Roman)
			if isSuccess {
				resultString := ToRoman(result)
				return resultString, true
			}
		} else if numTypeFirst == Arabic && numTypeTwo == Arabic {
			result, isSuccess := ConvertAndCalculate(numFirstString, numTwoString, operator, Arabic)
			if isSuccess {
				resultString := strconv.Itoa(result)
				return resultString, true
			}
		} else if (numTypeFirst == Roman && numTypeTwo == Arabic) || (numTypeFirst == Arabic && numTypeTwo == Roman) {
			panic(DIFFERENTNUMBERS)
			return "", false
		} else {
			panic(NOVALID)
			return "", false
		}

	} else {
		panic(FORMAT)
		return "", false
	}
	panic(NOOPERATOR)
	return "", false
}
