package main
import "fmt"


func main(){
  x:=42
  fmt.Println(x)
  fmt.Println(&x)
  y:=&x
  *y=56
  fmt.Println(x)
}