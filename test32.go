package main
import "os"
import "log"

func main(){
  f,err:=os.OpenFile("File1.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
  if err!=nil{
     log.Println(err)
  }
  defer f.Close()
  if _,err:=f.WriteString("Hello How are you\n");err!=nil{
     log.Fatal(err)
  }
}