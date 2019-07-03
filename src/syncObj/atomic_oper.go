package main   
import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int32 = 0
	wait_group := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wait_group.Add(1)
		go func() {
			atomic.AddInt32(&data, 1)
			wait_group.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wait_group.Add(1)
		go func() {
			atomic.AddInt32(&data, -1)
			wait_group.Done()
		}()
	}

	wait_group.Wait()
	fmt.Println(data)
}