package main
import "fmt"
import "bytes"
import "encoding/json"

type User struct{
  Name string
  Age int
}

func main(){
  user:=User{Name:"Bapan",Age:34}
  buf:=new(bytes.Buffer)
  encoder:=json.NewEncoder(buf)
  err:=encoder.Encode(user)
  if err!=nil{
     fmt.Println("Error Occured")
  }
  fmt.Println(buf)
}