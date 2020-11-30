package main
import "encoding/json"
import "fmt"
import "net/http"


type City struct{
  Name string  `json:"name"`
  Area uint64  `json:"area"`
}

func mainlogic(w http.ResponseWriter,r *http.Request){
  if r.Method=="POST"{
     var temp City
     decoder:=json.NewDecoder(r.Body)
     err:=decoder.Decode(&temp)
     fmt.Println(temp)
     if err!=nil{
        panic(err)
     }
     defer r.Body.Close()
     fmt.Printf("City is %s city and area is %d",temp.Name,temp.Area)
     w.WriteHeader(http.StatusOK)
     w.Write([]byte("201-Created"))
  }else{
     w.WriteHeader(http.StatusMethodNotAllowed)
     w.Write([]byte("405-Method not allowed"))
  }
}

func main(){
   http.HandleFunc("/city",mainlogic)
   http.ListenAndServe(":8000",nil)
}