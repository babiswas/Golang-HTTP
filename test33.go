package main
import "fmt"
import "os"
import "log"

func main(){
   path,err:=os.Executable()
   if err!=nil{
    log.Println(err)
   }
   fmt.Println(path)
}