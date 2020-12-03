package main
import "fmt"
import "io/ioutil"
import "os"


func main(){
   dir,err:=ioutil.TempDir(".","Test")
   if err!=nil{
     fmt.Println(err)
   }
   defer os.Remove(dir)
   fmt.Println(dir)
}