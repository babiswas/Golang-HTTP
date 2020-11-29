package main
import "fmt"
import "os"
import "strconv"


func main(){
  sum:=0
  for _,element:=range os.Args{
     num,_:=strconv.Atoi(element)
     sum=sum+num
  }
  fmt.Println(sum)
}