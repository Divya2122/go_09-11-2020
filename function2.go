package main 
  
import( 
    "fmt"
    "strings"
) 
  

func joinstr(element...string)string{ 
    return strings.Join(element, "-") 
} 
  
func main() { 
    
  
   fmt.Println(joinstr()) 
     
    
   fmt.Println(joinstr("divya")) 
   fmt.Println(joinstr("this", "is", "Divya")) 
   fmt.Println(joinstr("d", "i", "v", "y", "a")) 
     
} 