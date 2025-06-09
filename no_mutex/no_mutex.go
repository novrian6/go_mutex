package main

import (
	"fmt"
	"sync"
	"time"
)

var balance = 100

func withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	if balance >= amount {
		time.Sleep(100 * time.Millisecond) // Simulate delay
		balance -= amount
		fmt.Println("Withdrawn:", amount)
	} else {
		fmt.Println("Insufficient balance")
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(70, &wg)
	go withdraw(70, &wg)
	wg.Wait()
	fmt.Println("Final balance:", balance)
}
