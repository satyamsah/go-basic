

package main

import( "fmt")
 func main(){

var g1 geometry = circle{2.} 
fmt.Println("area circle ",g1.area())

var g2 geometry = rectangle{10,30} 
fmt.Println("area rectangle",g2.area())

}

type circle struct{
 	radius float64;
}

type rectangle struct{
 	length float64;
	width float64;
}

type geometry interface{
	area() float64 //method to be overrideen

}
//method overriding by rectangle class
func (rec rectangle) area() float64{
	return rec.length*rec.width
}
//method overriding by circle class
func (cir circle) area() float64{
	return 3.14*cir.radius*cir.radius
}

	