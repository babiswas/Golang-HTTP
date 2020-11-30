package main
import "fmt"

func func1(num int)func() int{
   return func()(int){
       return num+10
   }
}

func main(){
  m:=func1(4)
  fmt.Println(m())
  
}