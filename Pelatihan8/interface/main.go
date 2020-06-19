// memanggil method di func
// structnya jadi parameter

package main

import "fmt"

type dict interface {
	GetMorningGreeting() string
}

type englishDict struct{}

type koreaDict struct{}

func (en englishDict) GetMorningGreeting() string {
	return "Hello, good morning"
}

func (kr koreaDict) GetMorningGreeting() string {
	return "좋은 아침"
}

func main() {
	enDict := englishDict{}
	krDict := koreaDict{}
	printGreeting(enDict)
	printGreeting(krDict)
}

func printGreeting(d dict) {
	fmt.Println(d.GetMorningGreeting())
}

// func printEnglishGreeting(en englishDict) {
// 	fmt.Println(en.GetMorningGreeting())
// }

// func printKoreaGreeting(kr koreaDict) {
// 	fmt.Println(kr.GetMorningGreeting())
// }
