package main

import( "fmt")
 func main(){

   mp:= newMessagePrinter("hello world")
   fmt.Println(mp.message)	
}


// parameterized constructor

func newMessagePrinter(msg string) *messagePrinter{
	localconstruct:= new(messagePrinter);
	localconstruct.message =msg 
	return localconstruct
	
}

type messagePrinter struct{
 message string
}

