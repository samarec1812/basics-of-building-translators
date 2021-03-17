package main

import (
	"./calc"
	"./check"
	"./proc"
	"bufio"
	"fmt"
	"os"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

char arr[100];
int input() {

   fgets(arr, 100, stdin);
   for(int i = 0; i < strlen(arr); i++) {
     if (arr[i] == '0') {
         arr[i] = '\0';
			break;
     }
   }


return 0;
}
*/
import "C"

func main() {

	var answer string
	for {
		fmt.Println("Select the mode:")
		fmt.Println("1: Expression in a text file")
		fmt.Println("2: Expression in a console")
		fmt.Println("3: Exit")
		fmt.Print("Your select: ")
		_, _ = fmt.Scan(&answer)
		switch answer {
		case "1":
			file, err := os.Open("input.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			s := bufio.NewScanner(file)
			if s == nil {
				fmt.Println("Исходный файл пуст")
				return
			}
			for s.Scan() {
				expression := s.Text()
				expression = proc.Processing(expression)
				if check.CorrectString(expression) {
					fmt.Println("GOOD STRING")
					outString, arrOut := calc.PerformToPostfix(expression)
					fmt.Println(outString)
					if arrOut != nil {
						answ, err := calc.Evalation(arrOut)
						if err != nil {
							fmt.Println(err)
							continue
						}

						fmt.Println("answer is :", answ)

					} else {
						fmt.Println("answer is :", outString)
					}
				}
			}
		case "2":
			 //C.input();
			 ////fmt.Println(s, reflect.TypeOf(s))
             //input := C.GoString(*C.char(C.arr))
             //fmt.Println(input)

			myscanner := bufio.NewScanner(os.Stdin)
			myscanner.Scan()
			inputStr := myscanner.Text()

			inputStr = proc.Processing(inputStr)
		//	calc.FillDict(inputStr)
			if check.CorrectString(inputStr) {
				fmt.Println("GOOD STRING")
				outString, arrOut := calc.PerformToPostfix(inputStr)
				if outString != "" {
					fmt.Println(outString)
					if arrOut != nil {
						answ, err := calc.Evalation(arrOut)
						if err != nil {
							fmt.Println(err)
							continue
						}

						fmt.Println("answer is :", answ)
					} else {
						fmt.Println("answer is :", outString)
					}
				} else {
					continue
				}
			}
		case "3":
			fmt.Println("Close program")
			return
		default:
			fmt.Println("This answer is undefined. Choose more")
			fmt.Println()

		}
	}
	/*
		// fmt.Println(infixToPostfix("    1+2+3*7/5  "))
	    expr := "   a-  b-(  c*d  )/(c/e)   "
	    expr = proc.Processing(expr)
	    fmt.Println("Our string after processing", expr)
	    if check.CorrectString(expr)  {
	        fmt.Println("GOOD STRING")
	        outString := calc.PerformToPostfix(expr)
	        fmt.Println(outString)

	    } else {
	        fmt.Println("BAD STRING")
	    }
	   // fmt.Println(infixToPostfix("( A + B ) * C - ( D - E ) * ( F + G )"))
	*/
}
