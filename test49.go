package main
import "fmt"
import "os"
import "log"
import "io"
import "strings"

func main(){
  name:="Bapan Biswas"
  tpl:=`<h1>`+name+`</h1>`
  f,err:=os.Create("test1.html")
  if err!=nil{
   log.Fatal(err)
  }
  io.Copy(f,strings.NewReader(tpl))
  fmt.Println("Created html file")
}