package main
import "fmt"


type PyamentMethod interface{
	Pay(amount int) string
}


type PayMee struct {}
type ClickUp struct {}


func (p PayMee ) Pay(amount int) string {
	return "Done ! by paymee"
}

func (c ClickUp) Pay(amount int) string {
	return  "Done ! by click app"
}


func Checkout(p PyamentMethod, amount int){
	resutl := p.Pay(amount)
	fmt.Println("Payment",resutl)
}


func main() {
	fmt.Println("Interfaces in Go!")	

	payMee :=  PayMee{}
	ClickUp_ := ClickUp{}

	Checkout(payMee, 50000)
	Checkout(ClickUp_, 12000)
}