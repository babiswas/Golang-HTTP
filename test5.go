package main
import "fmt"
import "net/http"
import "github.com/julienschmidt/httprouter"
import "log"

func getstring(command string,arguments ...string)string{
   str1:=command + " "
   for i:=0;i<len(arguments);i++{
      str1=str1+arguments[i]+" "
   }
   return str1
}
func task1(w http.ResponseWriter,r *http.Request,params httprouter.Params){
   fmt.Fprintf(w,getstring("Command","test1","test2","test3"))
}

func task2(w http.ResponseWriter,r *http.Request,params httprouter.Params){
  fmt.Fprintf(w,getstring("command",params.ByName("name")))
}

func main(){
   router:=httprouter.New()
   router.GET("/task1",task1)
   router.GET("/task2:name",task2)
   log.Fatal(http.ListenAndServe(":8000",router))
}