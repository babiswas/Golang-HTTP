package main
import "fmt"
import "encoding/json"

type User struct{
  Name string `json:"name"`
  Adress string `json:"adress"`
}

func main(){
  user:=User{Name:"Bapan",Adress:"Delhi"}
  data,err:=json.Marshal(user)
  if err!=nil{
     fmt.Println("Error Occured")
  }
  fmt.Println(string(data))
  user1:=User{}
  err=json.Unmarshal(data,&user1)
  fmt.Println(user1)
}