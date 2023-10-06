package tugas

import "fmt"

// Silahkan kerjakan latihan pada soal berikut :

// Buatlah sebuah struct dengan nama Student. Struct ini mempunyai properti Name dan Class. Lalu, dia memiliki 2 method yaitu SetMyName(name string) dan CallMyName().

//     SetMyName(name string) harus bisa melakukan perubahan pada nama si student
//     CallMyName() akan menampilkan Hello, My Name is <nama kamu>.

type Student struct {
	Name  string
	Class int8
}

func (s *Student) SetMyName(name string) {
	s.Name = name
}

func (s *Student) CallMyName() {
	fmt.Println("Hello, my Name is", s.Name)
}

func Pointer() {
	student := Student{Name: "Kresna", Class: 10}
	fmt.Println(student)

	student.CallMyName()
	student.SetMyName("Rangga")
	student.CallMyName()
}
