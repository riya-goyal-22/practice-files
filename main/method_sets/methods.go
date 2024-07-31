package methods
import (
	"fmt"
	"time"
)

func funcAdd(a,b int)int{
	return a+b
}

func summation(a,b int)(int,interface{}){
	start:=time.Now()
	f:=funcAdd(a,b)
	duration:=time.Since(start)
	return f,duration
}

type dog struct{
	name string
}



func (d dog)walk(){
	d.name="xefgvgg"
	fmt.Println("my name is ",d.name,"and i am walking")
}

func (d *dog)run(){
	d.name="rover"
	fmt.Println("my name is ",d.name,"and i am running")
}

type animal interface{
	walk()
	run()
}

func animalRun(a animal){
	a.walk()
}

func change_name(d dog,s string)dog{
	d.name=s
	return d
}

func ChangeName(d *dog,s string)dog{
	d.name=s
	return *d
}

func Methods(){
	//fmt.Println(summation(3,4))
	d1:=dog {"henry"}
	fmt.Println(d1)
	d1.walk()
	fmt.Println(d1)
	// d1=change_name(d1,"jenny")
	// fmt.Println(d1)
	// d2:=&dog {"henry2"}
	// fmt.Println(*d2)
	// ChangeName(d2,"jenny2")
	// fmt.Println(*d2)
	// d1.walk()
	// d1.run()
	// animalRun(d1)
	//d2:=&dog{"bob"}

	// x:=2
	// fmt.Println(x)
	// fmt.Println(&x)
}
