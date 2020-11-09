package main
import "fmt"
function main(){
arr:[4]string{"this", "is", "the","divya"}
fmt.println("array",arr)
myslice :=arr[1:4]
fmt.println("slice:",myslice)
fmt printf("length of the slice:%d",len(slice))
fmt printf("capacity of the slice:%d",cap(myslice))
}