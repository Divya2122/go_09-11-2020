package main 
   
import "fmt"
   
func main() { 
   

    var a = 600
       
  
    var p = &a 
   
    fmt.Println("Value stored in a before changing = ", a) 
    fmt.Println("Address of a = ", &a) 
    fmt.Println("Value stored in pointer variable p = ", p) 
  
    
    fmt.Println("Value stored in y(*p) Before Changing = ", *p) 
  
    
    *p = 500 
  
     fmt.Println("Value stored in a(*p) after Changing = ",a) 
  
} 