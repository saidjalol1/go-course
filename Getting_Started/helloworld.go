package main
import "fmt"



func main(){
	fmt.Println("Hello, World!")


	// Variables
	// A variable is a storage location that has a name and a type. 
	// The value of a variable can be changed at runtime.
	var student1 string = "John";
	var student2 = "Jane";
	x := 10;
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)


	// constants
	// value must be assigned at compile time and cannot be changed at runtime
	// There are two types of constants in Go: typed and untyped constants
	// Typed constants:
	const PI float64 = 3.14
	fmt.Println(PI)
	// Untyped constants:
	const E = 2.71828
	fmt.Println(E)


	// Go output functions
	// fmt.Print() - prints the output without a newline
	// fmt.Println() - prints the output with a newline
	// fmt.Printf() - prints the output with formatting
	var name, surname string = "John", "Doe"
	fmt.Print("Hello, ", name, " ", surname, "\n")
	fmt.Println("Hello,", name, surname)
	fmt.Printf("Hello, %s %s\n", name, surname)


	// Go formatting verbs
	// The following verbs can be used to format all types of data in Go:
	// %v - the value in a default format
	// %T - a Go-syntax representation of the type of the value
	// %d - base 10 integer
	var integer_data = 42
	var text_data = "Hello, World!"
	fmt.Printf("Integer: %d\n", integer_data)
	fmt.Printf("Text: %s\n", text_data)
	fmt.Printf("Value: %v\n", integer_data)
	fmt.Printf("Type: %T\n", integer_data)
	

	// Go data types:
	// Go has three basic data types:
	// bool - represents a boolean value (true or false)
	// string - represents a sequence of characters
	// numeric types - represents numeric values (int, float, etc.)
	
	var age int = 30
	var height float64 = 5.9
	var isStudent bool = true
	var name2 string = "Alice"
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.1f\n", height)
	fmt.Printf("Is Student: %t\n", isStudent)
	fmt.Printf("Name: %s\n", name2)
	

}