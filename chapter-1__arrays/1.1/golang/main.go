package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ProductOfAllOtherElements(array []int) []int {
	if len(array) <= 1 {
		return []int{1}
	}

	// [Step] Prep array of left products
	leftProducts := make([]int, len(array))
	leftProducts[0] = array[0]
	for i := 1; i < len(array); i++ {
		leftProducts[i] = leftProducts[i-1] * array[i]
	}

	// [Step] Prep array of right products
	rightProducts := make([]int, len(array))
	rightProducts[len(array)-1] = array[len(array)-1]
	for i := len(array) - 2; i >= 0 ; i-- {
		rightProducts[i] = rightProducts[i+1] * array[i]
	}

	// [Step] Generate array of products except i
	products := make([]int, len(array))
	products[0] = rightProducts[1]
	products[len(array)-1] = leftProducts[len(array)-2]

	for i := 1; i < len(array)-1; i++ {
		products[i] = leftProducts[i-1] * rightProducts[i+1]
	}

	return products
}


func parseString(s string) ([]int, error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	file, err := os.Open("../input.txt")
	defer file.Close()
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		array, err := parseString(scanner.Text())

		if err != nil {
			break
		}
		
		fmt.Println(array)
		answer := ProductOfAllOtherElements(array)
		fmt.Println(answer)
		fmt.Println()
	}
}
