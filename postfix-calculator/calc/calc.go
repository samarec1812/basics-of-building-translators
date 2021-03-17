package calc

import (
	"../check"
	"fmt"
	"strconv"

	//"strconv"
	"strings"
)



func PriorityOperation(operation byte) int {
	switch operation {
	case '+':
		return 1
	case '-':
		return 1
	case '*':
		return 2
	case '/':
		return 2
	case '!':
		return 3
	}
	return 0
}


func PerformToPostfix(expr string) (string, []string) {
	ToPostfix := make([]string, 0, 0)
	if len(expr) == 1 {
		return expr, []string{expr}
	}
    operation := make([]byte, 0, 0)
    strNumber := ""
    fmt.Println(expr)
    for i := 0; i < len(expr); i++ {

    	if check.IsLetter(expr[i]) {
    		chInStr := string(expr[i])
    		ToPostfix = append(ToPostfix, chInStr)
		} else if check.IsDigit(expr[i]) {
			numPoint := 0
			for j := i; j < len(expr); j++ {
				if expr[j] == '.' { numPoint++ }
				if numPoint > 1 {
					fmt.Println("NUM IS NOT CORRECT")
				return "", []string{}
				}
				if !check.IsDigit(expr[j]) && expr[j] != '.'  {
					break
				}
				strNumber += string(expr[j])
				i = j
			}
				ToPostfix = append(ToPostfix, strNumber)
				strNumber = ""
				//			}
		} else if expr[i] == '-' && (i == 0 || expr[i-1] == '(' && (check.IsLetter(expr[i+1]) || check.IsDigit(expr[i+1]))) {
			//operation = append(operation, expr[i])
			str := string(expr[i])
			str += string(expr[i+1])
			ToPostfix = append(ToPostfix, str)
			i++
		} else if check.IsOperation(expr[i]) {
			if len(operation) == 0 {
				operation = append(operation, expr[i])

			} else if PriorityOperation(operation[len(operation)-1]) >= PriorityOperation(expr[i])  {
				for len(operation) > 0 && (PriorityOperation(operation[len(operation)-1]) >= PriorityOperation(expr[i])) {
					chInStr := string(operation[len(operation)-1])
					operation = operation[:len(operation)-1]
					ToPostfix = append(ToPostfix, chInStr)
				}
				operation = append(operation, expr[i])
			} else {
				operation = append(operation, expr[i])
			}
		} else if expr[i] == '(' {
			operation = append(operation, expr[i])

		} else if expr[i] == ')' {

			for operation[len(operation) - 1] != '(' {
				chInStr := string(operation[len(operation)-1])
				operation = operation[:len(operation)-1]
				ToPostfix = append(ToPostfix, chInStr)
			}
			operation = operation[:len(operation)-1]
		}
	}
	for len(operation) > 0 {
		chInStr := string(operation[len(operation)-1])
		operation = operation[:len(operation)-1]
		ToPostfix = append(ToPostfix, chInStr)
	}
	strToPostfix := strings.Join(ToPostfix, "")
	return strToPostfix, ToPostfix
}

func calculate(operation string, a float64, b float64) (float64, error) {
	if operation == "+" {
		return a + b, nil
	}
	if operation == "-" {
		return a - b, nil
	}
	if operation == "*" {
		return a * b, nil
	}
	if operation == "/" {
		if b == 0 {
			// fmt.Println("ERROR: Division by zero")
			return 0, fmt.Errorf("ERROR: Division by zero")
		}
			return a / b, nil
	}
	fmt.Println("ERROR: operation is not find")
	return 0, fmt.Errorf("ERROR: operation is not find")
}
var Dict map[string]float64
func FillDict(expr string) map[string]float64 {
	dict := make(map[string]float64)
	for _, symbol := range expr {
		_, ok := dict[string(symbol)]
		if check.IsLetter(byte(symbol)) && !ok {
			fmt.Printf("Enter value for %s: ", string(symbol))
			var numInStr string
			fmt.Scan(&numInStr)
			numInFloat, _ := strconv.ParseFloat(numInStr, 64)
			dict[string(symbol)] = numInFloat
		}
	}
    return dict
}
var DictNumber map[string]float64
/*func fillDictNumber(ToPostfix []string) map[string]float64 {
	dict := make(map[string]float64)
	for _, value := range ToPostfix {
		_, ok := dict[string(value)]
		if check.IsDigit2(value) && !ok {
			numInFloat, _ := strconv.ParseFloat(numInStr, 64)
			dict[string(symbol)] = numInFloat
		}
	}
}*/

func Permutation(s string) bool {
	if (check.IsLetter2(s) || check.IsDigit2(s))  { return true}
    return false
}

func Permutation2(s string, s1 string, s2 string) bool {
	if (check.IsLetter2(s) || check.IsDigit2(s)) && (check.IsLetter2(s1) || check.IsDigit2(s1)) && s2 == "-" { return true}
    return false
}
func Evalation(ToPostfix []string) (float64, error) {
	result := make([]float64, 0, 0)
	Dict := FillDict(strings.Join(ToPostfix, ""))

  //  DictNumber := make(map[float64]float64)
	for index := 0; index < len(ToPostfix); index++ {
		fmt.Println(result, ToPostfix)
		if check.IsLetter2(ToPostfix[index]) {
			result = append(result, Dict[ToPostfix[index]])
		} else if ToPostfix[index] == "!" {
			a := result[len(result)-1]
			result = result[:len(result)-1]
			result = append(result, -a)
		} else if len(ToPostfix[index]) > 1 && check.IsLetter(ToPostfix[index][1]) {
	        result = append(result, -Dict[string(ToPostfix[index][1])])
		} else if ToPostfix[index] == "+" ||
			ToPostfix[index] == "*" || ToPostfix[index] == "/" || (ToPostfix[index] == "-" && index != 0 &&
			(Permutation(ToPostfix[index-1]) || index == len(ToPostfix)-1 )) {

			a, b := result[len(result)-1], result[len(result)-2]
			result = result[:len(result)-2]
			c, err := calculate(ToPostfix[index], b, a)
			if err != nil {
				return 0, err
			}
			result = append(result, c)
			fmt.Println("calculate: ", result, c)
		} else if ToPostfix[index] == "-" && index != len(ToPostfix)-1 {
			fmt.Println(ToPostfix[index], index)

			if check.IsLetter2(ToPostfix[index+1]) && (index == 0) { /* || (index > 0 && (check.IsDigit2(ToPostfix[index-1])||
				check.IsLetter2(ToPostfix[index-1]) && check.IsOperation2(ToPostfix[index-2])))) { */
				a := Dict[ToPostfix[index+1]]
				Dict[ToPostfix[index+1]] = -a

			} else if check.IsDigit2(ToPostfix[index+1]) && (index == 0) { /* || (index > 0 && (check.IsDigit2(ToPostfix[index-1])||
				check.IsLetter2(ToPostfix[index-1]) && check.IsOperation2(ToPostfix[index-2])))) { */

				number, _ := strconv.ParseFloat(ToPostfix[index+1], 64)
				result = append(result, -number)
				index += 1
			}

		} else {
			num, err := strconv.ParseFloat(ToPostfix[index], 64)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, num)
		}
	}

	return result[len(result)-1], nil
}
