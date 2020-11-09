package main
  
import "fmt"
  

func swap(a, b int)int{
 
    var p int
    p= b
    b=a
    a=p
    
   return p 
}
  

func main() {
 var a int = 1.0
 var b int = 2.0
  fmt.Printf("a = %d and b = %d", a, b)
  
 
 swap(a, b)
   fmt.Printf("\na = %d and b = %d",a, b)
}