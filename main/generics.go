package main

import("fmt")

type numbers interface{
	~int |float64|string
}

func addT[T numbers](a, b T)T{
return a+b
}
type myalias int
func main(){
	var n myalias=25
	str1:="hii"
	str2:="hello"
	fmt.Println(addT(str1,str2))
	fmt.Println(addT(n,5))
	fmt.Println(addT(4.5,3.4))
}