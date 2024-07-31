package sort
import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Name < a[j].Name }

	

func Sort(){
	// a:=[]int{10}
	// s:=[]string{"hi","riya","tomorrow","leave","happy","music"}
	// sort.Strings(s)
	// fmt.Println(s)
	// a=append(a,2,9,6,8,4,6,2,7)
	// fmt.Println(a)
	// sort.Ints(a)
	// fmt.Println(a)
	people := []Person{
		{"Bobb", 99},
		{"Henry", 12},
		{"Tom", 17},
		{"Jenny", 22},
		{"Ivin",34},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// sort.Slice(people, func(i, j int) bool {
	// 	return people[i].Age < people[j].Age
	// })
	// fmt.Println(people)

	
	
}