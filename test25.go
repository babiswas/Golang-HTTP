package main
import "fmt"
import "os"
import "io"
import "strings"

func main(){
  f,err:=os.Create("File1.txt")
  if err!=nil{
    fmt.Println(err)
  }
  defer f.Close()
  reader:=strings.NewReader("Hello world How are you")
  io.Copy(f,reader)
}