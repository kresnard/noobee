package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

func main() {
	asynchronous()
	// synchronous()
}

func synchronous() {
	users, _ := readFile("./data.json")
	users = usdToIdr(users)
	log.Println(users[0])
}

func asynchronous() {
	userCh, _ := readFileConcurrent("./data.json")

	userWithIdr1 := usdToIdrConcurrent(userCh)
	userWithIdr2 := usdToIdrConcurrent(userCh)
	userWithIdr3 := usdToIdrConcurrent(userCh)
	userWithIdr4 := usdToIdrConcurrent(userCh)

	users := merge(userWithIdr1, userWithIdr2, userWithIdr3, userWithIdr4)

	counter := 0
	for user := range users {
		_ = user
		counter++
	}
	log.Println("counter", counter)

}

func readFile(filename string) (users []User, err error) {
	now := time.Now()
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	users = []User{}

	err = json.Unmarshal(dataByte, &users)
	if err != nil {
		return
	}

	log.Println("success read data in", time.Since(now).Microseconds(), "microseconds")

	return

}

func readFileConcurrent(filename string) (<-chan User, error) {
	userCh := make(chan User)
	now := time.Now()
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	users := []User{}

	err = json.Unmarshal(dataByte, &users)
	if err != nil {
		return nil, err
	}

	go func() {
		for _, user := range users {
			userCh <- user
		}
		close(userCh)
	}()

	log.Println("success read data in", time.Since(now).Microseconds(), "microseconds")

	return userCh, nil

}

func usdToIdr(users []User) (newUsers []User) {
	now := time.Now()
	for _, user := range users {
		time.Sleep(10 * time.Millisecond)
		user.Salary *= 100
		newUsers = append(newUsers, user)
	}
	log.Println("success change usd to idr", time.Since(now).Seconds(), "s")

	return
}

func usdToIdrConcurrent(user <-chan User) <-chan User {
	now := time.Now()
	userCh := make(chan User)

	go func() {
		for u := range user {
			time.Sleep(10 * time.Millisecond)
			newData := u
			newData.Salary = newData.Salary * 15_000
			userCh <- newData
		}
		close(userCh)
	}()

	log.Println("success change usd to idr in", time.Since(now).Nanoseconds(), "ns")

	return userCh
}

func merge(users ...<-chan User) <-chan User {
	usersCh := make(chan User)

	wg := sync.WaitGroup{}

	for _, user := range users {
		wg.Add(1)
		go func(user <-chan User) {
			for u := range user {
				usersCh <- u
			}
			wg.Done()
		}(user)
	}

	go func() {
		wg.Wait()
		close(usersCh)
	}()

	return usersCh
}
