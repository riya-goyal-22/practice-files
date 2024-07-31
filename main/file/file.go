package file

import(
	"fmt"
	// "io/ioutil"
	"io"
	"os"
	"log"
	"bytes"
)

func (p Person)writeOut(w io.Writer)error{
	_,err:=w.Write([]byte(p.name))
	return err
}

func FileCreate(){
	p:=Person{
		name:"mickey",
		age:30,
		place:"Disney",
	}
	f,err :=os.Create("demo2.txt")
	if err!=nil{
		log.Fatalf("error %s",err)
	}
	defer f.Close()
	var b bytes.Buffer
	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())

	// s:=[]byte("hello gophers")
	// _,err=f.Write(s)
	// if err!=nil{
	// 	log.Fatalf("error %s",err)
	// }
	// b:=bytes.NewBufferString("heloooooo")
	// fmt.Println(b.String())
	// b.WriteString(" how are you?")
	// fmt.Println(b.String())
	// b.Reset()
	// fmt.Println(b.String())
	// b.Write([]byte("this is tuesday"))
	// fmt.Println(b.String())

	//callback
	// fmt.Println(doOp(4,5,sub))

	fmt.Println(fact(6))
	xb,err:=readFile("demo2.txt")
	if err!=nil{
		log.Fatalf("readfile in main:%s\n",err)
	}
	fmt.Println("xb:",xb)
	fmt.Println("xb in string:",string(xb))
}

func readFile(fn string)([]byte,error){
	xb,err:=os.ReadFile(fn)
	if err!=nil{
		return nil,fmt.Errorf("error in readfile func :%s\n",err)
	}
	return xb,nil
}

func fact(a int)int{
	if a==1{
		return 1
	}
	return a*fact(a-1)
}

func doOp(a int ,b int,f func(int,int)int)int{
	return f(a,b)
}

func addition(a int ,b int)int{
	return a+b
}
func sub(a,b int)int{
	return a-b
}

