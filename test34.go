package main
import "fmt"
import "os"
import "log"

func main(){
   path,err:=os.Getwd()
   if err!=nil{
     log.Fatal(err)
   }
   fmt.Println(path)
}