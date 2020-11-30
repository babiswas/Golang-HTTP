package main
import "encoding/json"
import "fmt"
import "os"
import "strings"

func main(){
 type Person struct{
  Name string  `json:"name"`
  Hobby string `json:"hobby"`
  Age  int `json:"age"`
 }
 person:=Person{Name:"Bapan",Hobby:"Coding",Age:32}
 payload,err:=json.Marshal(person)
 if err!=nil{
    fmt.Println("Error Occured")
 }
 var temp Person
 fmt.Println(string(payload))
 encoder:=json.NewEncoder(os.Stdout)
 err=encoder.Encode(person)
 jreader:=strings.NewReader(string(payload))
 decoder:=json.NewDecoder(jreader)
 err=decoder.Decode(&temp)
 if err!=nil{
   fmt.Println("Error occured")
 }
 fmt.Println(temp)
}