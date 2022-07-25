package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	fmt.Println("running graceful shutdown sample")
	fmt.Println("enter singla os.Interrupt, os.Kill to shutdown")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		counter := 1
		for {
			select {
			case <-c.Done():
				fmt.Println("sub ctx done 1")
				return
			default:
				time.Sleep(1 * time.Second)
			}
			fmt.Printf("%04d : PING\n", counter)
			counter += 1
		}
	}(ctx)

	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		counter := 1
		for {
			select {
			case <-c.Done():
				fmt.Println("sub ctx done 2")
				return
			default:
				time.Sleep(1 * time.Second)
			}
			fmt.Printf("%04d : PONG\n", counter)
			counter += 1
		}
	}(ctx)

	fmt.Println("context wait start")
	<-ctx.Done()
	fmt.Println("context wait end")

	fmt.Println("wait group wait start")
	wg.Wait()
	fmt.Println("wait group wait end")

}
