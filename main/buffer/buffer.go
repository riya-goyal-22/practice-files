package buffer
import (
	"fmt"
	"bytes"
	
)

func Buffer(){
	var buf bytes.Buffer
	buf.WriteString("hello ")
	//buf.Reset()
	buf.WriteString("gophers!")
	fmt.Println(buf.String())
	
	
}