package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	reportLoaded     bool `default:false`
	diagnosticReport [][]bool
}

func (s *Submarine) CalculatePowerConsumption() int {
	var gammaRate, epsilonRate string
	for index, _ := range s.diagnosticReport[0] {
		if s.mostCommon(index, s.diagnosticReport) {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}
	gamma, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonRate, 2, 64)
	return int(gamma) * int(epsilon)
}

func (s *Submarine) filterEntries(report [][]bool, column int, value bool) (filteredReport [][]bool) {
	for _, row := range report {
		if row[column] == value {
			filteredReport = append(filteredReport, row)
		}
	}
	return filteredReport
}

func (s *Submarine) CalculateLifeSupportRating() int {
	var oxygenGeneratorRating, co2ScrubberRating string
	filteredReport := s.diagnosticReport

	for index, _ := range s.diagnosticReport[0] {
		filteredReport = s.filterEntries(filteredReport, index, s.mostCommon(index, filteredReport))
		if len(filteredReport) == 1 {
			break
		}
	}
	for _, value := range filteredReport[0] {
		if value {
			oxygenGeneratorRating += "1"
		} else {
			oxygenGeneratorRating += "0"
		}
	}
	filteredReport = s.diagnosticReport

	for index, _ := range s.diagnosticReport[0] {
		filteredReport = s.filterEntries(filteredReport, index, !s.mostCommon(index, filteredReport))
		if len(filteredReport) == 1 {
			break
		}
	}
	for _, value := range filteredReport[0] {
		if value {
			co2ScrubberRating += "1"
		} else {
			co2ScrubberRating += "0"
		}
	}

	oxygenGenerator, _ := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	co2Scrubber, _ := strconv.ParseInt(co2ScrubberRating, 2, 64)
	return int(oxygenGenerator) * int(co2Scrubber)
}

func (s *Submarine) mostCommon(column int, report [][]bool) bool {
	counter := 0
	for _, row := range report {
		if row[column] {
			counter += 1
		} else {
			counter -= 1
		}
	}
	return counter >= 0
}

func (s *Submarine) LoadDiagnosticReport(scanner *bufio.Scanner) {
	var parsedEntry []bool
	for scanner.Scan() {
		row := scanner.Text()
		parsedEntry = make([]bool, len(row))
		for index, char := range row {
			parsedEntry[index] = strings.Compare(string(char), "1") == 0
		}
		s.diagnosticReport = append(s.diagnosticReport, parsedEntry)
	}
	return
}

func main() {
	fmt.Println("day 3")
	inputFile, _ := os.Open("./day3/input")
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	s := Submarine{}
	s.LoadDiagnosticReport(scanner)
	fmt.Println(s.CalculatePowerConsumption())
	fmt.Println(s.CalculateLifeSupportRating())
}
