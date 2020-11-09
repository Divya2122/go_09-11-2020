package main 
  
import "fmt"
  
func main() { 
  
    
    var My_map = make(map[float64]string) 
    fmt.Println(My_map) 
  
    
    My_map[1.3] = "divya"
    My_map[1.5] = "chandu"
    fmt.Println(My_map) 
} 