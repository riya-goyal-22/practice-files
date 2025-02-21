package strings
import(
	"fmt"
	// "strconv"
	"strings"
)
func Strings(){
	var r rune='😁'
	fmt.Printf("%v\t %T\n",r,r)  //gives unicode value for 😁
	var placeOfInterest string= `❤️`

    fmt.Printf("plain string: ")
    fmt.Printf("%s", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("hex bytes: ")
	
	
    for i := 0; i < len(placeOfInterest); i++ {
        fmt.Printf("%x ", placeOfInterest[i])
    }
    fmt.Printf("\n")

	fmt.Println("-----------------------------")
	s := "We went to eat at multiple cafe"
cafe := "cafe"
if p := strings.Index(s, cafe); p != -1 {
    p += len(cafe)
    s = s[:p] + "s" + s[p:]
}
fmt.Println(s)

	var buffer [10]byte
	for _,val:=range buffer{
		fmt.Println(val)
	}
	
}