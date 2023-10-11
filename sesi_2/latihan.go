package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := map[string]string{
		"Name":    "NooBee",
		"Class":   "Backend Intermediate",
		"Address": "Jakarta",
	}

	// Key : Name, Value : NooBee
	// Key : Class, Value : Backend Intermediate
	// Key : Addres, Value : Jakarta

	// buatlah sebuah function print untuk nge handle hasil seperti dibawah
	// pastikan menggunakan goroutine, agar urutan hasilnya itu bisa berbeda beda

	// Key : Name, Value : NooBee
	// Key : Class, Value : Backend Intermediate
	// Key : Addres, Value : Jakarta

	wg := sync.WaitGroup{}
	for key, value := range arr {
		wg.Add(1)

		go func(key, value string) {
			print(key, value)
			wg.Done()
		}(key, value)
	}
	wg.Wait()
}

func print(key string, value string) {
	fmt.Printf("Key : %v, Value : %v \n", key, value)
}
