package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules := make([]rule, 0)
	readFile(&rules)
	result := 0

	for _, r := range rules {
		result += r.Test()
	}

	fmt.Println(result)
}

type rule struct {
	Min      int
	Max      int
	Letter   string
	Password string
}

func (r *rule) Test() int {
	c := 0
	for idx, l := range strings.Split(r.Password, "") {
		//fmt.Println(idx+1, l)
		if l == r.Letter && (idx+1 == r.Min || idx+1 == r.Max) {
			c++
		}
	}
	//fmt.Println(c)
	//fmt.Println("---")
	if c == 1 {
		return 1
	}
	return 0
}

func readFile(rules *[]rule) {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txt := scanner.Text()
		split1 := strings.Split(txt, ":")
		//fmt.Println(split1)
		password := strings.Trim(split1[1], " ")

		split2 := strings.Split(split1[0], " ")
		letter := split2[1]
		minmax := strings.Split(split2[0], "-")
		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])
		r := rule{
			Max:      max,
			Min:      min,
			Letter:   letter,
			Password: password,
		}
		*rules = append(*rules, r)
		//fmt.Println(r)

	}

}
