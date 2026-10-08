package main
import "fmt"

func main() {
	fmt.Println("Slices in Go!")


	/*  Slices are similar to arrays, but are more powerful and flexible.

		Like arrays, slices are also used to store multiple values of the same type in a single variable.

		However, unlike arrays, the length of a slice can grow and shrink as you see fit.

		In Go, there are several ways to create a slice:

			- Using the []datatype{values} format
			- Create a slice from an array
			- Using the make() function   */

	// 1. Using the []datatype{values} format
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []string{"Hello", "World", "!"}
	fmt.Println(slice1)
	fmt.Println(slice2)

	// 2. Create a slice from an array
	array := [5]int{1, 2, 3, 4, 5}
	slice3 := array[1:4]
	fmt.Println(slice3)


	// 3. Using the make() function
	// The make() function is used to create slices, maps, and channels in Go.
	// The make() function takes three arguments: the type of the slice, 
	// the length of the slice, and the capacity of the slice.
	slice4 := make([]int, 5, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // prints the length of slice4
	fmt.Println(cap(slice4)) // prints the capacity of slice4


	// Modifying slice elements
	slice1[0] = 10
	fmt.Println(slice1) // prints the modified slice1

	// Appending elements to a slice
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println(slice1) // prints the modified slice1

	// Length and capacity of a slice
	fmt.Println(len(slice1)) // prints the length of slice1
	fmt.Println(cap(slice1)) // prints the capacity of slice1

	// Copying a slice
	slice5 := make([]int, len(slice1))
	copy(slice5, slice1)
	fmt.Println(slice5) // prints the copied slice5
}
