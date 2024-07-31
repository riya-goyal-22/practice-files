package interface
import (
	"fmt"
	
)
type Reader interface{
	read()string
}
type Writer interface{
	write(s string)

}

type doc struct{
	name string
}

func (d doc)read()string{
	return d.name
}
func (d *doc)write(s string){
	d.name =s
}
type readwrite interface{
	Reader
	Writer
	read()string
}

type Floatish struct{
	f float64
}
type stringg struct{
s string
}

func (s stringg)String()string{
	fmt.Println("stringer function called")
	return s.s
}



func Interface(){
	f1:=Floatish{34.5}
	fmt.Println(f1)
	d1:=&doc{"hello.txt"}
	var r readwrite=d1
	r.write("sample")
	fmt.Println(r.(*doc).name)
	s1:=stringg{"heloooo"}
	fmt.Println(s1)
}