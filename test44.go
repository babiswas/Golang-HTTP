package main
import "fmt"
import "os"
import "log"


func main(){
  f,err:=os.Create("test1.txt")
  defer f.Close()
  if err!=nil{
     log.Fatal(err)
  }
  l,err:=f.WriteString("Hello World")
  if err!=nil{
     log.Fatal(err)
  }
  fmt.Println(l,"bytes written sucessfully") 
}