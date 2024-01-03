## Тестовая задача «Калькулятор»

Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b.

### Поддерживаемые операции

| Операция  | Символ  |
|-----------|---------|
| Сложение  | ```+``` |
| Вычитание | ```-``` |
| Умножение | ```*``` |
| Деление   | ```/``` |

### Поддерживаемые целые числа

Арабские: ```0, 1, 2, 3, 4, 5, 6, 7, 8, 9```

Римские: ```I, II, III, IV, V, VI, VII, VIII, IX```

### Требования к работе программы

Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.

Калькулятор умеет работать как с арабскими (1,2,3,4,5…), так и с римскими (I,II,III,IV,V…) числами.

Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.

Калькулятор умеет работать только с целыми числами.

Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен указать на ошибку и прекратить работу.

При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.

При вводе пользователем не подходящих чисел приложение выводит ошибку в терминал и завершает работу.

При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выводит ошибку и завершает работу.

Результатом операции деления является целое число, остаток отбрасывается.

Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна указать на исключение.

### Примеры
```txt
Input:
1 + 2

Output:
3

Input:
VI / III

Output:
II

Input:
I - II

Output:
В римской системе нет отрицательных чисел.

Input:
I + 1

Output:
Используются одновременно разные системы счисления.

Input:
1

Output:
Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).

Input:
1 + 2 + 3

Output:
Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).
```