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

	for idx := range inputData {

		q := target - idx
		for jdx := range inputData {
			if idx == jdx {
				continue
			}
			if tmp, ok := inputData[q-jdx]; ok {
				fmt.Println(idx, tmp, jdx, " = ", idx+tmp+jdx)
				result = idx * tmp * jdx
				break
			}
		}
		if result != 0 {
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
