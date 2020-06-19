// jalaninnya kalo di windows pake set
// kalo di linux pake export

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// type config struct {
// 	dbUser     string `json:db_user`
// 	dbPassword string `json:db_password`
// }

type appConfig map[string]interface{}

func main() {
	// config := os.Getenv("fileConfig")
	// fmt.Println("file Config: ", config)
	// fmt.Println("FOO:", os.Getenv("FOO"))
	//fmt.Println("File")
	//bacaFilePerBaris("file.txt")
	// tulisFile(config, []byte("Kalau berjalan\n"))
	// nambahTulisan(config, "prok...prok...prok\n")
	configFilePath := os.Getenv("configFile")
	appConfigContent, err := bacaFilePerBaris(configFilePath)

	if err != nil {
		panic("Read Config file failed")
	}

	var config = make(appConfig)
	configReader(config, appConfigContent, "=")
	fmt.Println(config)
}

func configReader(configFile appConfig, config []string, separator string) {
	for _, configVal := range config {
		c := strings.Split(configVal, separator)
		configFile[c[0]] = c[1]
	}
}

func bacaFilePerBaris(path string) ([]string, error) {
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
	temp := []string{}
	for s.Scan() {
		temp = append(temp, s.Text())
		// fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return temp, err
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
