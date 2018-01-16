package main
import ("fmt")
import("sync")
 
import "runtime"
func main(){
	runtime.GOMAXPROCS(8)
	var wg sync.WaitGroup
	wg.Add(2)
	ch:= make(chan int)
	go abcGen(ch,wg )
	go printer(ch,wg )
	fmt.Println("This comes first")
	//time.Sleep(100*time.Millisecond)
	defer wg.Wait()
	
	
}
func abcGen(ch chan int, wg sync.WaitGroup){
	defer wg.Done()
	for i:= 0; i<20;i++{
		//go fmt.Println(i)
		ch <- i
	}
	close(ch)
}

func printer(ch chan int,wg sync.WaitGroup){
	defer wg.Done()
	for l := range ch{
		fmt.Println(l)
	}

}