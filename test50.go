package main
import "os"
import "text/template"
import "log"


func main(){
  tpl,err:=template.ParseGlob("templates/*")
  if err!=nil{
    log.Fatal(err)
  }
  err=tpl.Execute(os.Stdout,nil)
}