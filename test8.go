package main
import "net/http"
import "fmt"



func middleware1(handler http.Handler)http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
     fmt.Println("Before processing request")
     handler.ServeHTTP(w,r)
     fmt.Println("After processing request")
  })
}

func mainlogic(w http.ResponseWriter,r *http.Request){
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Main logic Executed"))
}

func main(){
  handler:=http.HandlerFunc(mainlogic)
  http.Handle("/task",middleware1(handler))
  http.ListenAndServe(":8000",nil)
}