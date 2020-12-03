package main
import "fmt"
import "bufio"
import "log"
import "os"


func main(){
   f,err:=os.Open("File.txt")
   defer f.Close()
   if err!=nil{
     fmt.Println(err)
   }
   scanner:=bufio.NewScanner(f)
   for scanner.Scan(){
     fmt.Println(scanner.Text())
   }
   if err=scanner.Err();err!=nil{
      log.Fatal(err)
   }
}