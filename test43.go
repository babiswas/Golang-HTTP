package main
import "fmt"
import "flag"
import "bufio"
import "log"
import "os"


func main(){
  fptr:=flag.String("fpath","File1.txt","description")
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
  r:=bufio.NewReader(f)
  b:=make([]byte,3)
  for{
    n,err:=r.Read(b)
    if err!=nil{
      fmt.Println("Error reading the file:",err)
      break
    }
    fmt.Println(string(b[0:n]))
  }
}