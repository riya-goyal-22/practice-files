package main
import(
	"fmt"
)

type dog struct{
	name string
}

func (d dog) walk(){
	d.name="ruby"
	fmt.Println("I am  ",d.name," and I am walking")
}

func (d *dog) run(){
	d.name="ruby"
	fmt.Println("I am  ",d.name," and I am running")
}

func main(){
	d1:=dog{
		name:"shiro",
	}
	d2:=dog{
		name:"pluto",
	}
	d1.walk()
	fmt.Println(d1.name)
	d1.run()
	fmt.Println(d1.name)

	d2.walk()
	fmt.Println(d2.name)
	d2.run()
	fmt.Println(d2.name)

}