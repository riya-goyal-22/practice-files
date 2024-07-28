package main
 
import(
	"fmt"
	"os"
	"io"
)

func CreateFile(name string) *os.File{
	file,err:=os.Create(name+".txt")
	if err!=nil{
		panic(err)
	}
	return file
}

func WriteFile(f *os.File,con string){
	len,err:=io.WriteString(f,con)
	if err!=nil{
		panic(err)
	}
	fmt.Println("length:",len)
}

func ReadFile(filename string)[]byte{
	databytes,err:=os.ReadFile(filename)
	if err!=nil{
		panic(err)
	}
	return databytes
}

func main(){
	fmt.Println("::::::::Files in Golang:::::::::")
	content:="Shelly sells seashells on the sea shore. She sells seashells near the sea on sunday"
	var file_name string
	fmt.Print("Please enter your file name:")
	fmt.Scan(&file_name)
	file:=CreateFile(file_name)
	WriteFile(file,content)
	fmt.Println(string(ReadFile(file_name+".txt")))
	defer file.Close()
}