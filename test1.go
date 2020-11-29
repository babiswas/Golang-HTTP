package main
import "net/http"

func task(w http.ResponseWriter,r *http.Request){
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("hello"))
}

func main(){
  http.HandleFunc("/task",task)
  http.ListenAndServe(":8000",nil)
}

