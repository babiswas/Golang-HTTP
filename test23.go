package main
import "fmt"

func main(){
  type User struct{
    Name string
    Age int
  }
  type Person struct{
    Name string
    Age int
  }
  u:=User{Name:"Bapan",Age:23}
  p:=Person{Name:"Tapan",Age:32}
  var human interface{}=[2]interface{}{u,p}
  fmt.Println(human)
}


