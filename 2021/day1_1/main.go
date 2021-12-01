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
	prev := 0
	for idx := range input {
		if idx == 0 {
			prev = input[idx]
			continue
		}
		if input[idx] > prev {
			res++
		}
		prev = input[idx]
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
