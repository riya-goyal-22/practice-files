package pointerss

import(
	"fmt"
)

func derefrence(i *int)interface{}{
	return *i
}

func PassbyValue(arr [6]int){
	for i,val:=range arr{
		if val==3{
			arr[i]=33
		}
		
	}
}

func PassbyRef(arr *[6]int){
	arr[3]=9999
}


func Pointerss(){
	var i int=9
	var a =&i
	fmt.Println(a)
	fmt.Println(i)
	fmt.Println(derefrence(a))
	arr:=[6]int{3,4,5,6,7}
	PassbyValue(arr)
	fmt.Println(arr)
	PassbyRef(&arr)
	fmt.Println(arr)
	fmt.Println("--------------------------")
	slice:=make([]int,10,11)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice=append(slice,9,99,78,67)
	fmt.Println(arr[4])
	fmt.Println(slice[13])
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	

}