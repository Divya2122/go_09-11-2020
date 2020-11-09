package main 
  
import "fmt"
  
func main() { 
  
     
    var a int = 5748 
      
     
    var p *int
      
    
    p = &a
  
        
    fmt.Println("Value stored in a = ", a) 
    fmt.Println("Address of a = ", &a) 
    fmt.Println("Value stored in variable p = ", p) 
} 