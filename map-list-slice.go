package main

import( "fmt")
 func main(){

add := 2+1
add = add+2;
fmt.Println(add)
       mySlice := make([] float32,100)
	mySlice[0]=12.
	mySlice[1]=18.
	mySlice[34]=11.
	fmt.Println(mySlice)
	
	for idx,v := range mySlice{
	   println(idx,v)
	}
	
       myMap := make(map[int] string)
 fmt.Println(myMap)
       myMap[42] = "Foo"
       myMap[12] = "Bar"
       fmt.Println(myMap)
 fmt.Println(myMap[12])

for k,v := range(myMap){
	
	println(k,v)

}

foo :=1
  if(foo==1){
  	println("bar")
  }

switch foo{
 case 1:
    println("1")
 case 2:
    println("2")
}
	

}
	