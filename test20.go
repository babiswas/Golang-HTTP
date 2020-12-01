package main
import "encoding/json"
import "log"
import "net/http"
import "fmt"


type City struct{
  Name string
  Area uint64
}

func filterContentType(handler http.Handler)http.Handler{
   return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
      fmt.Println("Before executing middleware filter")
      if r.Header.Get("Content-type")!="application/json"{
         w.WriteHeader(http.StatusUnsupportedMediaType)
         w.Write([]byte("415-UnsupportedMediaType.Please send json"))
      }
      handler.ServeHTTP(w,r)
   })
}


func setServerTimeCookie(handler http.Handler)http.Handler{
   return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
      handler.ServeHTTP(w,r)
      cookie:=http.Cookie{Name:"Server-Time(UTC)"}
      http.SetCookie(w,&cookie)
      log.Println("Currently in the set server time middleware")
   })
}

func mainlogic(w http.ResponseWriter,r *http.Request){
   if r.Method=="POST"{
       var tempcity City
       decoder:=json.NewDecoder(r.Body)
       err:=decoder.Decode(&tempcity)
       if (err!=nil){
          panic(err)
       }
       defer r.Body.Close()
       log.Printf("City is %s and area is %d",tempcity.Name,tempcity.Area)
       w.WriteHeader(http.StatusOK)
       w.Write([]byte("201-Created"))
   }else{
     w.WriteHeader(http.StatusOK)
     w.Write([]byte("405-Method not allowed"))
   }
}

func main(){
   mainhandler:=http.HandlerFunc(mainlogic)
   http.Handle("/city",filterContentType(setServerTimeCookie(mainhandler)))
   http.ListenAndServe(":8000",nil) 
}





