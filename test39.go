package main
import "bufio"
import "fmt"
import "os"
import "log"
import "flag"


func main(){
  fptr:=flag.String("fpath","File1.txt","file path")
  flag.Parse()
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

