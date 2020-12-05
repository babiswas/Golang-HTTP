package main
import "bufio"
import "os"
import "log"
import "fmt"
import "flag"

func main(){
  fptr:=flag.String("name","test.txt","description")
  flag.Parse()
  fmt.Println(*fptr)
  f,err:=os.Open(*fptr)
  if err!=nil{
    log.Fatal(err)
  }
  defer func(){
     if err=f.Close();err!=nil{
        log.Fatal(err)
     }
  }()
  s:=bufio.NewScanner(f)
  for s.Scan(){
   fmt.Println(s.Text())
  }
  err=s.Err()
  if err!=nil{
     log.Fatal(err)
  }
}