//Viontina Dea IYP

package main

import (
	"fmt"
	"strings"
)

/*func match(pattern string, email string) {
	matched, _ := regexp.Match(pattern, []byte(email))
	if matched {
		fmt.Println("Email benar")
	} else {
		fmt.Println("Email tidak ada @")
	}
}*/

func main() {
	// input
	var email string
	fmt.Println("Email : ")
	fmt.Scanln(&email)

	//pattern := "^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$"
	// tidak boleh menggunakan kosong karena bisa berbahaya
	if email == "" || !strings.Contains(email, "@") {
		fmt.Println("Email tidak sesuai")
	} else if strings.HasPrefix(email, "@") {
		fmt.Println("Email tidak valid")
	} else if !strings.HasSuffix(email, ".com") {
		fmt.Println("Email tidak mengandung .com")
	} else if !strings.HasSuffix(email, "gmail") {
		fmt.Println("email tidak valid")
	} else {
		fmt.Println("Email sesuai")
		/*pattern := "[[:punct:]]{1}"
		match(pattern, email)*/
		//pattern := "com$"
		//match(pattern, email)
	}

}
