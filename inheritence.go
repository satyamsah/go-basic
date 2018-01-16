package main

import( "fmt")
 func main(){

   var mpe messagePrinterInterface = messagePrinter{"hi"}
   messagePrinterInterface.print()		
}



type messagePrinter struct{
 	message string
}

type messagePrinterInterface interface{
	print()

}

func (mp messagePrinter) print() {
	fmt.Println(mp.message)
}

