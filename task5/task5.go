package task5

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Task5(n int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)

	in := make(chan int)
	go func() {
		defer close(in)
		for {
			select {
			case <-ctx.Done():
				return
			case in <- rand.Intn(101):
			}
		}
	}()

	for v := range in {
		fmt.Println(v)
	}

	defer cancel()
}
