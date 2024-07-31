package structss
import(
	"fmt"
)

type person struct{
	sno int
	first string
	last string
}

func practi(P person) struct {
	sno   int
	first string
	last  string
	} {
return struct {
	sno   int
	first string
	last  string
}{
	sno:   1,
	first: P.first,
	last:  P.last,
}

	}

func Structs(){
	p1:=person{
		sno:1,
		first:"riya",
		last:"goyal",
	}
	var pracVar struct {
		sno   int
		first string
		last  string
	}
	pracVar = practi(p1)
	fmt.Println(pracVar)
}