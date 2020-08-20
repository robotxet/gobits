package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func f1(i int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1():", c1.Err())
		return
	case r := <-time.After(time.Duration(i) * time.Second):
		fmt.Println("f1():", r)
	}
	return
}

func f2(i int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(i)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Duration(i) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

func f3(i int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*i) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f2():", c3.Err())
		return
	case r := <-time.After(time.Duration(i) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("need a delay")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("delay:", delay)
	f1(delay)
	f2(delay)
	f3(delay)
}