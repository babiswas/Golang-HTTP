package main
import "fmt"
import "bytes"

func main(){
   var buf bytes.Buffer
   var buf1 bytes.Buffer
   fmt.Fprintf(&buf,"Hellow world")
   buf1.Write([]byte("Hellow world 100"))
   fmt.Println(buf.String())
   fmt.Println(buf1.String())
}