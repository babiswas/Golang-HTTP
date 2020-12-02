package main
import "fmt"

func main(){
  a:=make([]string,2,10)
  fmt.Println(cap(a))
  fmt.Println(len(a))
  b:=[]string{"Hello","bello"}
  c:=[]string{"Nello","Tello"}
  d:=[]string{"Bapan","Sapan","Tapan"}
  a=append(b,c...)
  a=append(a,d...)
  fmt.Println(a)
}