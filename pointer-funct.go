import( "fmt")
 func main(){
	result:= add(1,3,4)
        fmt.Println(result)
        var a int32 = 10
        var b int32 = 20
        swap(&a,&b)
        fmt.Println(a,b)
}

func swap(a *int32,b *int32 ){
   
   var temp int32= *a;
   *a=*b;
   *b=temp;
    
}

func add(terms ...int) int{
   result:= 0 
   for _,val:= range terms{
	result= result+val
   }
   return result;
}	
