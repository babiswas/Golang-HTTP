package main
import "strings"
import "log"
import "os"
import "io"

func main(){
   f,err:=os.Create("file.txt")
   if err!=nil{
       log.Fatal(err)
   }
   str1:="Hello world.I am Bapan"
   reader:=strings.NewReader(str1)
   io.Copy(f,reader)
}