package main
import (
	"fmt"
	"os"
	"runtime"
	"encoding/json"
	// "module/main/struct_methods"
	// "module/main/pointers"
	"module/main/new_make"
	// "github.com/riya-goyal-22/humans"
	// "module/calculator"
	// "dogs"
)
type Person struct{
	name string
	age int
	place string
}
func main(){
	// human.Speak()
	// human.Speaks()
	// human.NonHuman()
	// x:=calculator.Add(4,5)
	// fmt.Println(x)
	// dog.Bark()
	 

	//array
	// array:=[3]int{1,2,3}
	// fmt.Println(array)
	// var arr[5] string
	// for i:=0;i<4;i++{
	// 	fmt.Print("hello")
	// 	fmt.Scanln(&arr[i])
		
	// }
	// fmt.Println(arr)

	// p1:=Person{
	// 	name:"Riya",
	// 	age:21,
	// 	place:"Noida",
	// }

	// p1.speak()

	
	// FileCreate()
	


    gopath := os.Getenv("GOPATH")
    fmt.Println("GOPATH:", gopath)
	fmt.Println("GOROOT:", runtime.GOROOT())
	//Example()
	// structss.Structs()
	// pointerss
	dynamic.Dynamic()
	
}
func(p Person) speak(){
	fmt.Println("I am ",p.name)
}

type Internship struct{
	First string
	Last string
	College string
	Age int
}

func Example(){
	i1:=Internship{
		First:"riya",
		Last:"goyal",
		College:"JECRC",
		Age:20,
	}
	data,_:=json.Marshal(i1)
	fmt.Println(string(data))
}
