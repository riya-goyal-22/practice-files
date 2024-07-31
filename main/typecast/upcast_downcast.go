package typecast
import (
	"fmt"
)
type Animal interface{
	Speak()
}
type Cat struct{
	name string
}
type Dog struct{
	name string
}
func (c Cat)Speak(){
	fmt.Println("I'm a cat and my name is ",c.name)
}
func (d Dog)Speak(){
	fmt.Println("I'm a dog and my name is ",d.name)
}
func Typecast(){
	c1:=Cat{"milly"}
	fmt.Println(c1)
	d1:=Dog{"bob"}
	var a Animal
	a= d1
	if dog,ok:=a.(Dog);ok{
		fmt.Println(ok)
		fmt.Println(dog.name)
	}else{
		fmt.Println("a is not of type cat")
	}
	// if cat,ok:=a.(Cat);ok{
	// 	fmt.Println(ok)
	// 	fmt.Println(cat.name)
	// }else{
	// 	fmt.Println("a is not of type cat")
	// }
	cat,_:=a.(Cat)
	fmt.Println(cat.name)
	
}