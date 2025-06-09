package main

import (
	"fmt"
	"sync"
	"time"
)

var balance = 100
var mu sync.Mutex

func withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()         // Lock access to shared data
	defer mu.Unlock() // Ensure it's unlocked after this
	if balance >= amount {
		time.Sleep(100 * time.Millisecond)
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
