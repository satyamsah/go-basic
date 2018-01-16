package main

import( "fmt")
 func main(){
        foo:= newMyStruct()
		//foo:= new(myStruct)
	foo.myMap["Satyam"]=28
	fmt.Println(foo.myMap)
	
	
}

type myStruct struct{

    myMap mapofNameandAge[string] string

}

func newMyStruct() *myStruct{
	result:= new(myStruct) // point to the address like &mySruct
	//result1:= myStruct{} //points to the value of the address
	//result.myMap= make(map[string]string)
	result.myMap= map[string]string{}
	// return &result1
	return result
	
}