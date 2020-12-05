package main
import "io/ioutil"
import "fmt"
import "flag"
import "log"

func main(){
  fptr:=flag.String("name","test.txt","description")
  flag.Parse()
  data,err:=ioutil.ReadFile(*fptr)
  if err!=nil{
    log.Fatal(err)
  }
  fmt.Println(string(data))
}