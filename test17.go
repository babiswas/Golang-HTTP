package main
import "fmt"
import "io"
import "strings"

func main(){
   str1:="H"
   r:=strings.NewReader(str1)
   buf:=make([]byte,5)
   if _,err:=io.ReadFull(r,buf);err!=nil{
      fmt.Println(err)
   }
   fmt.Println(string(buf))
}