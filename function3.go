package main 
  
import "fmt"
  
   
 func GFG() func(i, j string) string{ 
     myf := func(i, j string)string{ 
          return i + j + "Divya"
     } 
    return myf 
 } 
    
func main() { 
    value := DIV() 
    fmt.Println(value("Welcome ", "to ")) 
} 