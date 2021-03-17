package check

import (
	"fmt"
	"strconv"
)

// Проверка, что символ является операцией
func IsOperation(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '!'
}

// Проверка, что символ является числом
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Проверка, что символ является строчной буквой латинского алфавита
func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z'
}

func IsLetter2(s string) bool {
	return "a" <= s && s <= "z"
}

func IsOperation2(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func IsDigit2(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil { return false }
	return true
}
/*
// Проверка на наличие инородных символов
func CheckSymbol(expr string) bool {
	if strings.ContainsAny(expr, ",.<>!@#$%^:&_=?\\`~№\";[]{}") {
		fmt.Println("String is contain incorrect symbol")
		return false
	}
	//if strings.ContainsRune()
	return true
}
*/
func isBracket(ch byte) bool {
	if ch == '(' || ch == ')' {
		return true
	}
    return false
}
// Проверка строку на инородные символы
func CheckSymbol(expr string) bool {
	flag := true
	for index, _ := range expr {
		if !IsDigit(expr[index]) && !IsLetter(expr[index])  && !IsOperation(expr[index]) && !isBracket(expr[index]) {
			flag = false
			break
		}
	}
    return flag
}

func CorrectString(expr string) bool {
	balance := 0
	if expr == "" {
		fmt.Println("ERROR: Empty input string")
		return false
	}
	if expr[0] == '+' || expr[0] == '*' || expr[0] == '/' {
		fmt.Println("ERROR: First symbol is binary operand")
		return false
	}
	if len(expr) == 1 && IsOperation(expr[0]) {
		fmt.Println("ERROR: string length is 1 and first symbol is operation")
		return false
	}
    if expr[0] == '(' {
    	balance += 1
	}
	if expr[0] == ')' {
		fmt.Println("ERROR: first symbol is right bracket")
		return false
	}
	//if !CheckSymbol(expr) {
	//	fmt.Println("ERROR: String is contain incorrect symbol")
	//	return false
	//}


	for i := 1; i < len(expr)-1; i++ {

		switch expr[i] {
		case '(':
			balance += 1
//			fmt.Println("balance now and left bracket: ", balance)
			if !IsOperation(expr[i-1]) && expr[i-1] == ')' {
				fmt.Println("1 error")
				return false
			}
			if !IsLetter(expr[i+1]) && !IsDigit(expr[i+1]) && expr[i+1] == ')' {
				fmt.Println("Error 1.1")
				return false
			}
		case ')':
			if balance == 1 {
				balance--
			} else {
				fmt.Println("2 error")
				return false
			}
			if (!IsLetter(expr[i-1]) && !IsDigit(expr[i-1])) || expr[i-1] == ')' {
				fmt.Println("Error 2.2")
				return false
			}
			if !IsOperation(expr[i+1]) && expr[i+1] == '(' {
				fmt.Println("Error 2.3")
				return false
			}
//			fmt.Println("balance now and right bracket: ", balance)
		case '+':
			if !IsLetter(expr[i-1]) && expr[i-1] == '(' && !IsDigit(expr[i-1]) {
				fmt.Println("Error 3")
				return false
			}
			if !IsLetter(expr[i+1]) && expr[i+1] == ')' && !IsDigit(expr[i+1]) {
				fmt.Println("Error 4")
				return false
			}
		case '-':
			if IsOperation(expr[i-1]) {
			// !IsLetter(expr[i-1]) && !IsDigit(expr[i-1]) && expr[i-1] != '('  {
				fmt.Println("Error 5")
				return false
			}
			if !IsLetter(expr[i+1]) && expr[i+1] == ')' && !IsDigit(expr[i+1]) {

				fmt.Println("Error 6")
				return false
			}
		case '*':
			if !IsLetter(expr[i-1]) && expr[i-1] == '(' && !IsDigit(expr[i-1]) {
				fmt.Println("Error 7")
				return false
			}
			if !IsLetter(expr[i+1]) && expr[i+1] == ')' && !IsDigit(expr[i+1]) {
				fmt.Println("Error 8")
				return false
			}
		case '/':
			if !IsLetter(expr[i-1]) && expr[i-1] == '(' && !IsDigit(expr[i-1]) {
				fmt.Println("Error 9")
				return false
			}
			if !IsLetter(expr[i+1]) && expr[i+1] == ')' && !IsDigit(expr[i+1]) {
				fmt.Println("Error 10")
				return false
			}
		}
		if IsLetter(expr[i]) && (IsLetter(expr[i-1]) || IsLetter(expr[i+1])) {
			fmt.Println("ERROR: some variables stand side by side ")
			return false
		}
		if IsLetter(expr[i]) && (IsDigit(expr[i-1]) || IsDigit(expr[i+1])) {
			fmt.Println("ERROR: variables and numbers stand round ")
			return false
		}
		if IsDigit(expr[i]) && (IsLetter(expr[i-1]) || IsLetter(expr[i+1])) {
			fmt.Println("ERROR: number and variables stand around ")
			return false
		}

  	}
	if !IsLetter(expr[len(expr) - 1]) && !IsDigit(expr[len(expr) -1]) && expr[len(expr)-1] != ')' {
		fmt.Println("Error 11 " + string(expr[len(expr) -1]))
		return false
	}
	if expr[len(expr) - 1] == ')' {
		balance--
	}
	if balance != 0 {
		fmt.Println("Error 12")
		return false
	}

	return true
}




