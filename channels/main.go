package main

import (
	"fmt"
	"time"
)


func addTwoNums(a int , b int , channelData chan int) {
	sum := a + b

	channelData <- sum // passing result to the channel
}

func streamSend(str string, channelD chan string){
	for _, val := range str{
		channelD <- string(val)
	}
	close(channelD)
}


func StreamReceive(channelD chan string){
	for val := range channelD {
		fmt.Print(val) 
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

func main(){
	fmt.Println("Channels in the Go !")

	
	// channelIntegers := make(chan int)
	// go addTwoNums(12, 45, channelIntegers)
	// resul := <- channelIntegers
	// println("Result: ", resul)


	// Stream 
	streamChan := make(chan string)
	go streamSend("Hallo !!!", streamChan)

	StreamReceive(streamChan)


	// buffered channel
	bufferedChan := make(chan int, 3)
	bufferedChan <- 12
	bufferedChan <- 45
	bufferedChan <- 78

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

}