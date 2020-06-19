// panic jarang dipakai, lebih sering dipakai di http server
// lebih sering pakai error yang biasa daripada panic
// kalau panic ngga perlu return error
// panic itu tetep return nilainya setelah menampilkan error sebelum exit

package main

import "fmt"

func main() {
	totalGaji := salaryCalculation()
	fmt.Println("Total Gaji Bulan ini", totalGaji)
}

func salaryCalculation() float64 {
	var totalGajiTemp float64

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	for i := 0; i <= 10; i++ {
		if i == 3 {
			panic("Fail Database")
		}
		totalGajiTemp++
	}
	return totalGajiTemp
}

// package main

// import "fmt"

// func main() {
// 	totalGaji, err := salaryCalculation()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Total Gaji Bulan ini", totalGaji)
// }

// type appError struct {
// 	customMessage string
// }

// func (a appError) Error() string {
// 	return a.customMessage
// }

// func newAppError(message string) *appError {
// 	return &appError{customMessage: message}
// }

// func salaryCalculation() (t float64, e error) {
// 	var totalGajiTemp float64
// 	var aErr error = nil
// 	defer func() {
// 		if err := recover(); err != nil {
// 			e = newAppError(fmt.Sprintf("%v", err))
// 			t = totalGajiTemp
// 		}
// 	}()
// 	for i := 0; i <= 10; i++ {
// 		if i == 3 {
// 			panic("Fail")
// 		}
// 		totalGajiTemp++
// 	}
// 	return totalGajiTemp, aErr
// }
