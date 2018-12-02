package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	dataString := string(data)
	eachNumber := strings.Split(dataString, "\n")

	var frequency = 0
	var mapOfFrequencies = map[int]bool{
		frequency: true,
	}
	var repeatanceFound = false

	for j := 0; !repeatanceFound; j++ {
		for _, number := range eachNumber {
			i, _ := strconv.Atoi(number)
			frequency += i
			if _, ok := mapOfFrequencies[frequency]; ok && !repeatanceFound {
				fmt.Printf("Frequency %d was repeated\n", frequency)
				repeatanceFound = true
			} else {
				mapOfFrequencies[frequency] = true
			}
		}
	}

	fmt.Printf("Final frequency: %d\n", frequency)

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
