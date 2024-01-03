package internal

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var romanLettersMap = map[string]int{
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}
var arabicLettersMap = map[int]string{
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	5:   "V",
	4:   "IV",
	1:   "I",
}

func ReadConsole() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line = strings.TrimSpace(line)
	line = strings.ToUpper(line)
	line = strings.ReplaceAll(line, " ", "")
	return line
}

func ValidateFormula(line string) bool {
	pattern := `^\w+[\-\+\/\*]\w+$`
	isValidFormula, _ := regexp.MatchString(pattern, line)
	return isValidFormula
}

func SplitByOperator(line string, operator string) (string, string) {
	words := strings.Split(line, operator)
	return words[0], words[1]
}

func FindOperation(line string) (string, string) {
	var operator string
	operators := [...]string{"+", "-", "*", "/"}

	for i := 0; i < len(operators); i++ {
		isOperator := strings.Contains(line, operators[i])
		if isOperator {
			operator = operators[i]
			break
		}
	}
	return line, operator
}

func WhatIsNumberType(line string) NumberType {
	pattern := `^[IVXLC]+$`
	isValidRoman, _ := regexp.MatchString(pattern, line)
	if isValidRoman {
		return Roman
	} else {
		_, err := strconv.Atoi(line)
		if err == nil {
			return Arabic
		} else {
			return Notype
		}
	}
}

func ConvertStringToArabic(line string, nType NumberType) int {
	numInt := 0
	switch nType {
	case Arabic:
		numInt, _ = strconv.Atoi(line)
	case Roman:
		numInt = ToArabic(line)
	default:
	}
	return numInt
}

func ToArabic(line string) int {
	numInt := 0
	for k := range line {
		if k < len(line)-1 && romanLettersMap[line[k:k+1]] < romanLettersMap[line[k+1:k+2]] {
			numInt -= romanLettersMap[line[k:k+1]]
		} else {
			numInt += romanLettersMap[line[k:k+1]]
		}
	}
	return numInt
}

func ToRoman(num int) string {
	result := ""

	keysArabicLettersMap := make([]int, 0)
	for k, _ := range arabicLettersMap {
		keysArabicLettersMap = append(keysArabicLettersMap, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keysArabicLettersMap)))

OuterLoop:
	for num != 0 {
		for _, k := range keysArabicLettersMap {
			if num >= k {
				result += arabicLettersMap[k]
				num -= k
				continue OuterLoop
			}
		}
	}
	return result
}

func CheckRangeNum(num int) bool {
	if num >= 1 && num <= 10 {
		return true
	} else {
		return false
	}
}

func ConvertAndCalculate(numFirstString string, numTwoString string, operator string, nType NumberType) (int, bool) {
	numFirstInt := ConvertStringToArabic(numFirstString, nType)
	numTwoInt := ConvertStringToArabic(numTwoString, nType)

	isValidRangeFirst := CheckRangeNum(numFirstInt)
	isValidRangeTwo := CheckRangeNum(numTwoInt)

	if isValidRangeFirst && isValidRangeTwo {
		result := Calculate(numFirstInt, numTwoInt, operator)
		if nType == Roman && result < 0 {
			panic(NEGATIVENUMBER)
			return 0, false
		} else {
			return result, true
		}
	} else {
		panic(NOVALIDRANGE)
		return 0, false
	}
}

func Calculate(numFirstInt int, numTwoInt int, operator string) int {
	result := 0
	switch operator {
	case "+":
		result = numFirstInt + numTwoInt
	case "-":
		result = numFirstInt - numTwoInt
	case "*":
		result = numFirstInt * numTwoInt
	case "/":
		result = numFirstInt / numTwoInt
	}
	return result
}
