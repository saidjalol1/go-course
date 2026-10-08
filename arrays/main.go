package main
import "fmt"

func main(){
	fmt.Println("Arrays in Go!")

	/*  Arrays are used to store multiple values of the same type in a single variable, 
	instead of declaring separate variables for each value.   */

	// There are two ways to declare an array in Go:
	// 1. Using the var keyword with fixed size
	var array1  = [5]int{1,2,3,4,5}
	var array2  = [3]string{"Hello", "World", "!"}
	fmt.Println(array1)
	fmt.Println(array2)

	// 2. Using the shorthand syntax with inferred size without using the var keyword
	array3 := [...]float64{1.1, 2.2, 3.3, 4.4}
	array4 := [...]float64{1.1, 2.2, 3.3, 4.4}
	fmt.Println(array3)			
	fmt.Println(array4)

	// Accessing array elements
	fmt.Println(array1[0]) // prints the first element of array1
	fmt.Println(array2[1]) // prints the second element of array2

	// Modifying array elements
	array1[0] = 10
	fmt.Println(array1) // prints the modified array1

	// Length of an array
	fmt.Println(len(array1)) // prints the length of array1
}