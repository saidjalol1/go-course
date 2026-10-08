package main

import (
	"fmt"
	"errors"
)


func divide_numbers(num int, num2 int ) (int, error){
	if num2 == 0{
		return  0 , errors.New("Cannot divide by zero")
	}

	return  num / num2, nil
}


func main(){
	fmt.Println("Error handling in Go!")

	result, err := divide_numbers(12, 12)
	if err != nil {
		fmt.Println("Error occured ", err)
		return
	}

	fmt.Println("Result:", result)

}