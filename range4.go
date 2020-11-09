package main 
  
import "fmt"
  
 
func main() { 
  
   
    student_rank_map := map[string]int{"divya": 3, 
                           "chandu": 2, "gopi": 1} 
  
    
    for student := range student_rank_map { 
        fmt.Println("Rank of", student, "is: ", 
                     student_rank_map[student]) 
    } 
  
    for student, rank := range student_rank_map { 
        fmt.Println("Rank of", student, "is: ", rank) 
    } 
} 