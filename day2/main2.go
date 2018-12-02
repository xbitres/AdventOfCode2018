package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	dataString := string(data)
	eachString := strings.Split(dataString, "\n")

	for nr, str := range eachString {
		for _, str2 := range eachString[nr+1:] {
			var nrCharDif = 0
			for i := 0; i < len(str); i++ {
				if str[i] != str2[i] {
					nrCharDif++
				}
			}

			if nrCharDif == 1 {
				fmt.Println(str)
				fmt.Println(str2)
			}
		}
	}
}
