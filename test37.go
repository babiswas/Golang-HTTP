package main
import "io/ioutil"
import "log"
import "os"
import "fmt"

func main(){
  file,err:=ioutil.TempFile(".","Test")
  if err!=nil{
    log.Fatal(err)
  }
  defer os.Remove(file.Name())
  fmt.Println(file.Name())
}