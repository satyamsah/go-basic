package main

import (

"io/ioutil"
"fmt"
"bufio"
"strings"
"os"
)


func readfile(filename string) map[string]int{
	wordcountMap:= make(map[string]int)

	inputfile, err:= os.Open(filename)
	if err!=nil{

		return wordcountMap
	}
	filescanner:= bufio.NewScanner(inputfile)
	for filescanner.Scan(){
	line:= strings.Split(strings.TrimSpace(string(filescanner.Text()))," ")
	   for _, word := range line{
			val,key := wordcountMap[word]
			if key{

			wordcountMap[word]=val+1
			}else {
			wordcountMap[word]=1
			}
		}
	}

	//for key,val:= range wordcountMap{
	//
	//fmt.Println (key,"      ",val)
	//}
	return wordcountMap
}

func main()  {

	files,err:= ioutil.ReadDir("/home/heavyvm/airavata/tools/go-tools/cis/files")
	if err!=nil{
		fmt.Print(err)
	}
	count:= 0
	numberoffiles:= len(files)
	fmt.Println( numberoffiles)
	wordcountfinalMap:= make(map[string]int)
	for _, singlefile:= range files{
		internalwordcountMap:= readfile("/home/heavyvm/airavata/tools/go-tools/cis/files/"+singlefile.Name())

		for key,val := range internalwordcountMap{
			mainval,iskeypresent := wordcountfinalMap[key]
			if iskeypresent{
				wordcountfinalMap[key]=mainval+val
			}else{
				wordcountfinalMap[key]=val
			}
		}

	}
	fmt.Println("key-- value")
	//for k,v := range  wordcountfinalMap{
	//	fmt.Println(k,"-",v)
	//}


}
