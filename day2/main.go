package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	horizontalPosition, depth, aim int
}

func (s *Submarine) move(x, y int) {
	s.horizontalPosition += x
	s.depth += y
}

func (s *Submarine) getCourse() int {
	return s.horizontalPosition * s.depth
}

func pilotSubmarine(s *Submarine, instructionsFile string) int {
	inputFile, _ := os.Open(instructionsFile)
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		x, y := 0, 0
		instruction := strings.Split(scanner.Text(), " ")
		direction := instruction[0]
		amount, _ := strconv.Atoi(instruction[1])
		switch direction {
		case "up":
			y = amount * -1
		case "down":
			y = amount
		case "forward":
			x = amount
		}
		s.move(x, y)

	}
	return s.getCourse()
}

func (s *Submarine) down(value int) {
	s.aim += value

}
func (s *Submarine) up(value int) {
	s.aim -= value

}
func (s *Submarine) forward(value int) {
	s.horizontalPosition += value
	s.depth += s.aim * value
}

func pilotSubmarineCorrectly(s *Submarine, instructionsFile string) int {
	inputFile, _ := os.Open(instructionsFile)
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		command := instruction[0]
		value, _ := strconv.Atoi(instruction[1])
		switch command {
		case "up":
			s.up(value)
		case "down":
			s.down(value)
		case "forward":
			s.forward(value)
		}
	}
	return s.getCourse()
}

func main() {
	fmt.Println("day 2")
	s1 := Submarine{}
	fmt.Println(pilotSubmarine(&s1, "./day2/input"))
	s2 := Submarine{}
	fmt.Println(pilotSubmarineCorrectly(&s2, "./day2/input"))
}
