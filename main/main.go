package main
import (
	"fmt"
	"os"
	"runtime"
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

	
}
func(p Person) speak(){
	fmt.Println("I am ",p.name)
}