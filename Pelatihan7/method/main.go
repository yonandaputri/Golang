package main

import (
	"fmt"
	"strconv"
)

//method itu fungsi yang menempel pada struct/type

type student struct {
	firstName, lastName string
	age                 int
}

// spasinya masuk parameter
// bebas student itu struct
func (bebas student) fullName(spasi string) string {
	umur := strconv.Itoa(bebas.age)
	return bebas.firstName + spasi + bebas.lastName + ", Umur : " + umur
}

func main() {
	student1 := student{
		firstName: "budi",
		lastName:  "anduk",
		age:       18,
	}
	var namaLengkap = student1.fullName(" ")
	fmt.Println(namaLengkap)
}
