package main

import "fmt"

type student struct {
	firstName, lastName string
	age                 int
	studentAddress      address
}

type address struct {
	homeAddress map[string]string
	postCode    string
}

func main() {
	newStudent := student{
		firstName: "budi",
		lastName:  "anduk",
		age:       18,
		studentAddress: address{
			homeAddress: map[string]string{"1": "subang", "2": "bandung"},
			postCode:    "41254",
		},
	}
	// & untuk ambil alamat
	// karena alamat memorinya sama jadi yg di newstudent juga berubah
	// kalau newstudent2 := newstudent alamat memorinya beda
	newStudent2 := &newStudent
	newStudent2.firstName = "tono"
	newStudent3 := newStudent2
	newStudent3.firstName = "ucup"
	newStudent3.lastName = "markucup"
	// * untuk ambil value
	//fmt.Println(newStudent)
	fmt.Println(newStudent.studentAddress.homeAddress["1"])
}

/*
student1 := student{
	firstName: "budi",
	lastName:  "anduk",
	age:       18,
}
// untuk ganti value pada pointer bisa pake * atau bisa langsung akses sendiri
student2 := &student1
(*student2).firstName = "TONO"
fmt.Println(student2)


student2.firstName = "TINI"
fmt.Println(student2)
*/

/*
// addressnya pake map
 // custom type
type student struct {
	firstName, lastName string
	age                 int
	address // kalau nama struct sama dengan nama variabel tidak perlu ditulis variabel lagi
}
type addressMap map[string]string

type address struct {
	homeAddress   addressMap
	officeAddress addressMap
}

func main() {
	student1 := student{
		firstName: "budi",
		lastName:  "anduk",
		age:       18,
		address: address{
			homeAddress:   addressMap{"alamat1": "subang", "alamat2": "bandung"},
			officeAddress: addressMap{"alamat1": "Bekasi", "alamat2": "Surabaya", "alamat3": "Semarang"},
		},
	}

	fmt.Println(student1)
}
*/
