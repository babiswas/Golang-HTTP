package main
import "fmt"

type User struct{
  Name string
  Age int
}

func (u *User)setname(name string){
   u.Name=name
}

func (u *User)setage(a int){
  u.Age=a
}

func (u *User)setUser(a User){
   u.Name=a.Name
   u.Age=a.Age
}

func main(){
  u:=User{Name:"Bapan",Age:15}
  fmt.Println(u)
  u.setname("Tapan")
  u.setage(34)
  fmt.Println(u)
  a:=User{Name:"Rajib",Age:32}
  fmt.Println(a)
  u.setUser(a)
  fmt.Println(u)
}