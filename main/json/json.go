package json
import (
	"fmt"
	"encoding/json"
)

type person struct{
	First string
	Last string
	Age int
}

func JSON(){
	p1:=person{
		First:"mickey",
		Last:"mouse",
		Age:30,
	}
	p2:=person{
		First:"minnie",
		Last:"mouse",
		Age:31,
	}
	people:=[]person{p1,p2}
	data:=marshall(people)
	fmt.Println(string(data))

	var people1 []person
	unmarshal(data,people1)
}

func marshall(people []person)[]byte{
	fmt.Println(people)
	data,err:=json.Marshal(people)
	if err!=nil{
		panic(err)
	}
	return data
}

func unmarshal(data []byte,people1 []person){
	err:=json.Unmarshal(data,&people1)
	if err!=nil{
		panic(err)
	}
	fmt.Println(people1)
	for i,v:=range people1{
		fmt.Println("person no ",i+1)
		var x=fmt.Println(v.First,v.Last,v.Age)
		
	}
}