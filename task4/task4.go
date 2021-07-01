package task4

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Task4(input <-chan interface{}, workersNum int) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	out := make(chan interface{})

	go func() {
		for v := range out {
			fmt.Println(v)
		}
	}()

	for i := 0; i < workersNum; i++ {
		go func(ctx context.Context) {
			for v := range input {
				// switch v.(type) {
				// case string:
				// 	cancel()
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

	return ctx, cancel
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
	}()

	ctx, cancel := Task4(in, 5)

	select {
	case <-sChan:
		fmt.Println("Cancel order")
		cancel()
	case <-ctx.Done():
		fmt.Println("Error: string in data stream")
	}
}
