package main
import "fmt"


type Toy interface{
  getcolor()
}

type Redtoy struct{
  Color string
  Id int
}

type Bluetoy struct{
  Color string
  Id int
}

func (r Redtoy)getcolor(){
  fmt.Println("Color is ",r.Color,r.Id)
}

func (b Bluetoy)getcolor(){
  fmt.Println("Color is",b.Color,b.Id)
}

func toycolor(b []Toy){
  for _,v:=range b{
     v.getcolor()
  }
}

func main(){
  r1:=Redtoy{Color:"Red1",Id:1}
  r2:=Redtoy{Color:"Red2",Id:2}
  r3:=Redtoy{Color:"Red3",Id:3}
  b1:=Bluetoy{Color:"Blue1",Id:4}
  b2:=Bluetoy{Color:"Blue2",Id:5}
  m:=[]Toy{r1,r2,r3,b1,b2}
  toycolor(m)
}