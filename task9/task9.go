package task9

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func numberGenerator(done <-chan interface{}) <-chan int {
	numbers := make(chan int)
	go func() {
		defer func() {
			fmt.Println("Input closed")
			close(numbers)
		}()
		for i := 1; i <= 25; i++ {
			select {
			case numbers <- i:
			case <-done:
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return numbers
}

func numberTransformer(done <-chan interface{}, input <-chan int) <-chan int {
	processed := make(chan int)

	go func() {
		defer func() {
			fmt.Println("Out closed")
			close(processed)
		}()
		for n := range input {
			select {
			case processed <- n * 2:
			case <-done:
				return
			}
		}
	}()

	return processed
}

func Task9() {
	done := make(chan interface{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigChan
		close(done)
		fmt.Println("Done closed! Canel pipeline...")
		time.Sleep(2)
	}()

	numbersIN := numberGenerator(done)
	numbersOUT := numberTransformer(done, numbersIN)

	for n := range numbersOUT {
		fmt.Println(n)
	}

}
