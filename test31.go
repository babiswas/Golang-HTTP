package main
import "fmt"
import "os"
import "bufio"
import "log"

func main(){
  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
     fmt.Println(scanner.Text())
  }
  if err:=scanner.Err();err!=nil{
     log.Fatal(err)
  }
}