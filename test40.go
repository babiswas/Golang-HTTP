package main
import "io/ioutil"
import "fmt"
import "flag"


func main(){
  fptr:=flag.String("name","test.txt","description")
  flag.Parse()
  fmt.Println(*fptr)
  data,err:=ioutil.ReadFile(*fptr)
  if err!=nil{
    fmt.Println(err)
    return
  }
  fmt.Println(string(data))
}
