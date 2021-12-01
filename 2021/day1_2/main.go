package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readFile()

	res := 0
	sum := 0
	for idx := range input {
		if idx < 3 {
			sum += input[idx]
			continue
		}
		if sum+input[idx]-input[idx-3] > sum {
			res++
		}
		sum = sum + input[idx] - input[idx-3]
	}

	fmt.Println("result:", res)
}

func readFile() []int {
	var inputData []int
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if number, err := strconv.Atoi(scanner.Text()); err == nil {
			inputData = append(inputData, number)
		}
	}
	return inputData
}
