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

	var twiceRep = 0
	var thirdRep = 0

	for _, str := range eachString {
		var occuranceMap = make(map[rune]int)
		for _, char := range str {
			occuranceMap[char]++
		}

		var twoFound, threeFound bool

		for _, nrRep := range occuranceMap {
			if nrRep == 2 && !twoFound {
				twoFound = true
				twiceRep++
			}
			if nrRep == 3 && !threeFound {
				threeFound = true
				thirdRep++
			}
		}
	}

	fmt.Printf("Checksum: %d * %d = %d", twiceRep, thirdRep, twiceRep*thirdRep)
}
