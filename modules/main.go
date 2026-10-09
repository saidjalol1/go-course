package main
import (
	"fmt"
	"github.com/saidjalol1/modules/utils"
)


func main(){
	fmt.Println("Structuring project files and packages in Go !")
	fmt.Println(utils.ExportedFunc())
	// fmt.Println(utils.unexportedFunc()) // This will cause a compile error since it's not exported

}