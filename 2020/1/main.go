package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputData := make([]int, 0)
	result := 0
	readFile(&inputData)

	sort.Ints(inputData)
	fmt.Printf("%#v\n", inputData)
	for idx := range inputData {
		//fmt.Println(idx)
		for jdx := len(inputData) - 1; jdx > idx && inputData[jdx]+inputData[idx] >= 2020; jdx-- {
			//fmt.Println(inputData[jdx], inputData[idx])
			if inputData[jdx]+inputData[idx] == 2020 {
				result = inputData[jdx] * inputData[idx]
				fmt.Println(inputData[jdx], inputData[idx])
			}
		}
	}

	fmt.Println("result:", result)
}

func readFile(inputData *[]int) {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if number, err := strconv.Atoi(scanner.Text()); err == nil {
			*inputData = append(*inputData, number)
		}

	}

}
