package main
import "fmt"


func main(){
	fmt.Println("Type assertion in Go !")

	var box any  = "Message !"

	str, ok := box.(string)

	if ok {
		fmt.Println("Success , It is a string", str)
	}else{
		fmt.Print("It is not a string")
	}
}