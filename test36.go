package main
import "fmt"
import "io/ioutil"


func main(){
   f,err:=ioutil.ReadDir(".")
   if err!=nil{
     fmt.Println(err)
   }
   for _,f:=range f{
      fmt.Println(f.Name())
   }
}