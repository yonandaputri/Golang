package main

import "fmt"

type gaji struct {
	basic, tunjangan, lembur float64
}

type karyawan struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []gaji
}

func (e *karyawan) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.basic)
		fmt.Println(info.tunjangan)
		fmt.Println(info.lembur)
	}
	fmt.Print("Total gaji Perbulan : ", e.totalGajiPerBulan())
	return "----------------------"
}

func (e *karyawan) totalSalary() float64 {
	var totalGaji float64
	for _, salary := range e.MonthlySalary {
		totalGaji += salary.basic + salary.tunjangan + salary.lembur
	}
	return totalGaji
}

func (e *karyawan) totalGajiPerBulan() {
	var gajiPerBulan float64
	for _, salary := range e.MonthlySalary {
		gajiPerBulan = salary.basic + salary.tunjangan + salary.lembur
	}
}

func main() {

	e := karyawan{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []gaji{
			gaji{
				basic:     15000.00,
				tunjangan: 5000.00,
				lembur:    2000.00,
			},
			gaji{
				basic:     16000.00,
				tunjangan: 5000.00,
				lembur:    2100.00,
			},
			gaji{
				basic:     17000.00,
				tunjangan: 5000.00,
				lembur:    2200.00,
			},
		},
	}

	fmt.Println(e.EmpInfo())
	fmt.Println(e.totalSalary())
	//fmt.Println(e.MonthlySalary[0].basic)

}
