package  main

import "fmt"
import (

	"time"

	"sync"
	"runtime"
)

func  main()  {

	runtime.GOMAXPROCS(2)
	var waitgrp sync.WaitGroup
	waitgrp.Add(2)
	go func(){
		defer waitgrp.Done()
		fmt.Print("Hi")
		time.Sleep(5)
	}()

	go func() {
		defer waitgrp.Done()
		fmt.Println("Hello")
	}()

	waitgrp.Wait()
}