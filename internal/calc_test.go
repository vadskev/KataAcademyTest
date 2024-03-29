package internal

import "testing"

func TestFromArabic(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"I+III", "IV"},
		{"IV-II", "II"},
		{"IX/IV", "II"},
		{"V*VI", "XXX"},
		{"3+1", "4"},
		{"6-4", "2"},
		{"4/3", "1"},
		{"9*4", "36"},
		{"I-V", "Числа выходят из диапазона от 1 до 10."},
		{"V+XI", "Числа выходят из диапазона от 1 до 10."},
		{"IV/X", "Числа выходят из диапазона от 1 до 10."},
		{"2-9", "-7"},
		{"9+9", "18"},
		{"3/0", "Числа выходят из диапазона от 1 до 10."},
		{"10*2", "20"},
		{"V+IV+", "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."},
		{"V$IV", "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."},
		{"V+IV+V", "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."},
	}

	for _, test := range tests {
		output, _ := Calc(test.input)

		if output != test.expected {
			t.Errorf("Expected '%s' to be '%s' but got '%s'", test.input, test.expected, output)
		}
	}
}
