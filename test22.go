package main
import "fmt"
func test(m ...interface{}){
   for i:=0;i<len(m);i++{
       fmt.Println(m[i])
   }
}

func main(){
  type A struct{
    Name string
    b int
  }
  a:=A{Name:"Bapan",b:4}
  type C struct{
     Name string
     c int
  }
  b:=C{Name:"Biswas",c:23}
  test(a,b)
}