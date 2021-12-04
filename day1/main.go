package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func counter(fileName string, slidingWindowSize int) (counter int) {
	var measurements []int = make([]int, int(slidingWindowSize))
	inputFile, _ := os.Open(fileName)
	populateSlidingWindow := true
	var pointer, previousTotal, currentTotal int
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		if populateSlidingWindow {
			measurements[pointer], _ = strconv.Atoi(scanner.Text())
			previousTotal += measurements[pointer]
			pointer += 1
			if pointer == slidingWindowSize {
				populateSlidingWindow = false
				pointer = 0
			}
			continue
		}
		for i, value := range measurements {
			if i == pointer {
				continue
			}
			currentTotal += value
		}
		newMeasurement, _ := strconv.Atoi(scanner.Text())
		currentTotal += newMeasurement

		measurements[pointer] = newMeasurement
		pointer = (pointer + 1) % slidingWindowSize

		if currentTotal > previousTotal {
			counter += 1
		}
		previousTotal = currentTotal
		currentTotal = 0
	}
	return
}

func main() {
	fmt.Println("day 1")
	fmt.Println(counter("./day1/input", 1))
	fmt.Println(counter("./day1/input", 3))
}
