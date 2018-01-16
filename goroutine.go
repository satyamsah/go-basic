package main
import ("fmt")
//import ("time")
import "runtime"
func main(){
	runtime.GOMAXPROCS(8)
	ch:= make(chan int)
	doneCh:= make(chan bool)
	go abcGen(ch)
	go printer(ch, doneCh)
	fmt.Println("This comes first")
	//time.Sleep(100*time.Millisecond)
	
	<- doneCh
	
}
func abcGen(ch chan int){
	for i:= 0; i<20;i++{
		//go fmt.Println(i)
		ch <- i
	}
	close(ch)
}

func printer(ch chan int, doneCh chan bool){
	for l := range ch{
		fmt.Println(l)
	}

	doneCh <- true
}