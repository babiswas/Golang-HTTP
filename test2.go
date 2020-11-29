package main
import "net/http"
import "fmt"
import "log"
import "github.com/gorilla/mux"
import "time"


func ArticleHandler(w http.ResponseWriter,r *http.Request){
  vars:=mux.Vars(r)
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w,"Category is:%v",vars["category"])
  fmt.Fprintf(w,"ID is %v",vars["id"])
}

func main(){
  r:=mux.NewRouter()
  r.HandleFunc("/article/{category}/{id:[0-9]+}",ArticleHandler)
  srv:=&http.Server{
    Handler:r,
    Addr:"127.0.0.1:8000",
    WriteTimeout:15*time.Second,
    ReadTimeout:15*time.Second,
  }
  log.Fatal(srv.ListenAndServe())
}