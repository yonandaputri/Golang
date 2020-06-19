// package io/ioutil buat data yang kecil, dari file besar dipecah-pecah
// package io/outil ngga bisa append, cuma bisa baca sama tulis
// package os untuk file yang besar lalu dipecah
// flag, environment variable, dan os argument (os.Arg) untuk nerima data dari luar
// lebih prefer pake environmet variable daripada flag dan os.Arg
// kalau pake flag informasi bisa kebongkar semua kaya password dan username
// go mux, go chi, go echo, go RM itu framework golang untuk REST API
// mutex : 1 variable itu bisa diupdate oleh banyak function yg jalan bersamaan
// mutex = locking, biar tidak ada variable yang ngambil kondisi terakhir yang sama, istilahnya rebutan isi yang sama
// flag, environment variable, dan os argument untuk naik ke production supaya kodingan bisa diaplikasikan

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//Ini juga penting untuk baca os environment variabel
	// nerima dari luar
	// otomatis bikin file .txt
	// Ini os environment variable yang banyak dipakai
	dbServer := os.Getenv("dbServer")
	dbUser := os.Getenv("dbUser")
	// config := os.Getenv("fileConfig")
	// ditampung slice of string
	// config := os.Args
	// fmt.Println("File Config:")
	// fmt.Println(config[1])
	// fmt.Println(config[2])
	// untuk naik ke production tidak perlu ganti ganti IP
	// [1] yang dibelakang untuk ambil hasil splitnya yang berupa IP
	// dbServer := strings.Split(config[1], "=")[1]
	// dbUser := strings.Split(config[2], "=")[1]

	fmt.Println("Server database kita ada di ", dbServer, " dengan user ", dbUser)
	for {
		fmt.Println("")
	}
	// fmt.Println("FOO:", os.Getenv("FOO"))
	//fmt.Println("File")
	//bacaFilePerBaris("file.txt")
	// tulisFile(config, []byte("Kalau berjalan\n"))
	// nambahTulisan(config, "prok...prok...prok\n")
}

func bacaFilePerBaris(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
func bacaFile(path string) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(string(body))
}

func tulisFile(fileName string, data []byte) {
	err := ioutil.WriteFile(fileName, data, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func nambahTulisan(fileName string, tulisan string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()

	if _, err = f.WriteString(tulisan); err != nil {
		log.Fatalf(err.Error())
	}

}
