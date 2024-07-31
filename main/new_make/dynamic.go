package dynamic
import(
"fmt"
)
func Dynamic(){
	p:=new(string)
	fmt.Println(*p)

	q:=make([][]int,3)
	fmt.Println(q)
}