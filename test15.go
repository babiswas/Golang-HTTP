package main
import "bufio"
import "strconv"
import "os"
import "fmt"

func main(){
  scanner:=bufio.NewScanner(os.Stdin)
  var a int64
  var b int64
  fmt.Println("Enter the value of a")
  scanner.Scan()
  a,_=strconv.ParseInt(scanner.Text(),10,64)
  fmt.Println("Enter the value of b")
  scanner.Scan()
  b,_=strconv.ParseInt(scanner.Text(),10,64)
  fmt.Printf("The sum of the number is %d",a+b)
}