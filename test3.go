package main
import "net/http"
import "fmt"
import "github.com/gorilla/mux"

func task1(w http.ResponseWriter,r *http.Request){
  vars:=mux.Vars(r)
  fmt.Println(vars["id"])
  w.Write([]byte("task1 executing"))
  w.WriteHeader(http.StatusOK)
}


func task2(w http.ResponseWriter,r *http.Request){
  vars:=mux.Vars(r)
  fmt.Println(vars["id"])
  w.Write([]byte("task2 executing"))
  w.WriteHeader(http.StatusOK)
}


func main(){
  r:=mux.NewRouter().StrictSlash(false)
  s:=r.PathPrefix("/task").Subrouter()
  s.HandleFunc("/",task1)
  s.HandleFunc("/{id}/details",task2)
  http.ListenAndServe(":8000",r) 
}

