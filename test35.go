package main
import "fmt"
import "os"
import "io/ioutil"
import "log"

func main(){
  f,err:=os.Open("File1.txt")
  defer f.Close()
  if err!=nil{
     fmt.Println(err)
  }
  buf,err:=ioutil.ReadAll(f)
  if err!=nil{
    log.Fatal(buf)
  }
  fmt.Println(string(buf))
}