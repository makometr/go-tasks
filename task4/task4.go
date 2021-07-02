package task4

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Task4(ctx context.Context, input <-chan interface{}, workersNum int) <-chan interface{} {
	out := make(chan interface{})
	isDone := make(chan interface{})

	go func() {
		for v := range out {
			fmt.Println(v)
		}
		isDone <- struct{}{}
	}()

	var wg sync.WaitGroup
	wg.Add(workersNum)

	for i := 0; i < workersNum; i++ {
		go func(ctx context.Context) {
			defer wg.Done()
			for v := range input {
				// switch v.(type) {
				// case string:
				// 	cancel() // bad practice?
				// 	return
				// }
				select {
				case <-ctx.Done():
					return
				case out <- v:
				}
			}
		}(ctx)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return isDone
}

// copy to main
func main() {
	sChan := make(chan os.Signal, 1)
	signal.Notify(sChan, syscall.SIGTERM, syscall.SIGINT) // ERROR
	defer close(sChan)

	in := make(chan interface{})
	go func() {
		for i := 0; i < 20; i++ {
			in <- i
			time.Sleep(100 * time.Millisecond)
		}
		in <- "error value"
		for i := 0; i < 20; i++ {
			in <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(in)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// isDone := task4.Task4(ctx, in, 5)
	isDone := Task4(ctx, in, 5)

	select {
	case <-sChan:
		fmt.Println("Cancel order")
	case <-isDone:
		fmt.Println("Datastream end!")
	}
}
