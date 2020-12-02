package main
import "fmt"
import "bytes"

func main(){
  var buf bytes.Buffer
  buf.WriteString("Hello World")
  fmt.Println(buf.String())
}