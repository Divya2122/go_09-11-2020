package main 
  
import "fmt"
  
func main() { 
  
     
    var map_1 map[int]int
  
    
    if map_1 == nil { 
      
        fmt.Println("True") 
    } else { 
      
        fmt.Println("False") 
    } 
  

    map_2 := map[int]string{ 
      
            90: "Divya", 
            91: "chandu", 
            92: "mamatha", 
            93: "babu", 
            94: "gopi", 
    } 
      
    fmt.Println("Map-2: ", map_2) 
} 
