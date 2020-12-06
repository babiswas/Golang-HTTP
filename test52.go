package main
import "fmt"

func findtype(i interface{}){
  switch i.(type){
    case string:
       fmt.Println("String block executed")
    case int:
       fmt.Println("Int block executed")
    default:
       fmt.Println("Default block executed")
  }
}

func main(){
   findtype("Naveen")
   findtype(77)
   findtype(22.23)
}