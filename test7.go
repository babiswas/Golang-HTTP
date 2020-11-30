package main
import "fmt"
import "github.com/gorilla/mux"
import "net/http"

func task1(w http.ResponseWriter,r *http.Request){
   param:=r.URL.Query()
   w.WriteHeader(http.StatusOK)
   fmt.Fprintf(w,"Got parameter %s",param["id"][0])
   fmt.Fprintf(w,"Got parameter %s",param["category"][0]) 
}

func main(){
  r:=mux.NewRouter()
  r.HandleFunc("/articles",task1)
  r.Queries("id","category")
  http.ListenAndServe(":7000",r)
}