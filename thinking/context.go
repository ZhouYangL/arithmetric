package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	go handle(ctx)
	time.Sleep(2*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
		
	}

}

func handle(ctx context.Context)  {
	select {
	case <- ctx.Done():
		fmt.Println("handle", ctx.Err())
	}
}


// 1 1     +
// 2 1 2 - +
// (1 （ 4 5 2  3 ） （ 6 8 ） + + + - + +
