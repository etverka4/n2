package main
import (
    "fmt"
    "time"
    "sync"
)
func main() {
    var swg sync.WaitGroup
    swg.Add(2)
    go_1 := func(){
        defer swg.Done()
        time.Sleep(1 * time.Second)
        fmt.Printf("done \n")
    }
    for i := 0; i < 2; i++{
        go go_1()
    }
    swg.Wait()
}
