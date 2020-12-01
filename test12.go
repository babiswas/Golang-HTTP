package main
import "fmt"
import "encoding/json"
import "strings"
import "os"
import "bytes"


type Person struct{
   Name string
   Age int
   Adress string
}

type User struct{
  Name string
  Age int
}


func main(){
  p:=Person{Name:"Bapan",Age:24,Adress:"Assam"}
  encoder:=json.NewEncoder(os.Stdout)
  err:=encoder.Encode(p)
  data,_:=json.Marshal(p)
  var temp Person
  reader:=strings.NewReader(string(data))
  decoder:=json.NewDecoder(reader)
  err=decoder.Decode(&temp)
  fmt.Println(temp)
  var buf bytes.Buffer
  user:=User{Name:"Tapan",Age:34}
  encoder=json.NewEncoder(&buf)
  err=encoder.Encode(user)
  if err!=nil{
    fmt.Println("Error Occured")
  }
  fmt.Println(buf)
}