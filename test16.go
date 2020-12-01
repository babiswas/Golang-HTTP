package main
import "fmt"
import "strings"
import "io"

func main(){
   r:=strings.NewReader("Hello World")
   buf:=make([]byte,2)
   for{
      n,err:=r.Read(buf)
      fmt.Println(n)
      fmt.Println(string(buf))
      if err==io.EOF{
         break
      }
   }
}