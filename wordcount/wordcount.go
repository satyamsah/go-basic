package main

import (

//"io/ioutil"
"fmt"
"bufio"
"os"
"strings"
)

func main()  {

	wordcountMap:= make(map[string]int)
	inputfile, err:= os.Open("/home/heavyvm/airavata/tools/go-tools/cis/files/moby-000.txt")

	if err!=nil{
		fmt.Println(err,"Fine not found")
	}



	filescanner:= bufio.NewScanner(inputfile)
	for filescanner.Scan(){
		line:= strings.Split(strings.TrimSpace(string(filescanner.Text()))," ")
		for _, word := range line{
			val,ifpresentkey := wordcountMap[word]
			if ifpresentkey{
				wordcountMap[word]=val+1
			}else {
				wordcountMap[word]=1
			}
		}
	}
	fmt.Println ("-key-val")
	for key,val:= range wordcountMap{

		fmt.Println (key,"      ",val)
	}

}
