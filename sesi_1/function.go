package main

import "fmt"

func FunctionMessage() {
	var car = make(map[string]string)
	car["name"] = "BMW"
	car["color"] = "Black"

	// alur
	// simpan hasil dari return function kedalam sebuah variable message
	message := returnString(car)
	// tampilkan hasil dari variable message
	showResult(message)

	// output => Mobil BMW berwarna Black

}

// buat 2 buah fungsi :
// 1 => fungsi yang mengembalikan sebuah string
// pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black

func returnString(car map[string]string) (sentence string) {
	return "Mobil " + car["name"] + " berwarna " + car["color"]
}

// 2 => fungsi yang menampilkan hasil dari kembalian string
// fungsi ini hanya bertugas untuk menampilkan kata
func showResult(message string) {
	fmt.Println(message)
}
