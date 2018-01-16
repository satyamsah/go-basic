package main

import( "fmt")
 func main(){
        foo := new(myStruct)
	foo.name = "bar"
	foo.age = 21
	fmt.Println(foo)
	
	foo1:= myStruct{"bar",21}
	foo1.name = "bar"
	foo1.age = 21
	fmt.Println(foo1)
	fmt.Print(foo1.name," ",foo1.age)
	
	
}

type myStruct struct{

    name string
    age  int32

}

