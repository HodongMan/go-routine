package main
 
import (
	"fmt"
	"time"
)
 
func hello() {
    fmt.Println("hello")
}
 
func main() {
	go hello()
	time.Sleep(time.Second * 1)
}