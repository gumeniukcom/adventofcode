package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputData := make(map[int]int, 0)
	result := 0
	readFile(inputData)

	const target = 2020

	//sort.Ints(inputData)
	//fmt.Printf("%#v\n", inputData)
	for idx := range inputData {
		if _, ok := inputData[target-idx]; ok {
			fmt.Println(idx)
			result = idx * (target - idx)
			break
		}

	}

	fmt.Println("result:", result)
}

func readFile(inputData map[int]int) {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if number, err := strconv.Atoi(scanner.Text()); err == nil {
			inputData[number] = number
		}

	}

}
