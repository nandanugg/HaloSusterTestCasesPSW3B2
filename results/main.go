package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the input file as an argument.")
		return
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var checks, httpReqFailed, dataReceived, dataSent string
	var avg, min, med, max, p90, p95 string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "checks") {
			checks = extractPercentage(line)
		} else if strings.Contains(line, "http_req_failed") {
			httpReqFailed = extractPercentage(line)
		} else if strings.Contains(line, "data_received") {
			dataReceived = extractData(line)
		} else if strings.Contains(line, "data_sent") {
			dataSent = extractData(line)
		} else if strings.Contains(line, "{ expected_response:true }") {
			avg = extractValue(line, "avg")
			min = extractValue(line, "min")
			med = extractValue(line, "med")
			max = extractValue(line, "max")
			p90 = extractValue(line, "p\\(90\\)")
			p95 = extractValue(line, "p\\(95\\)")
		}
	}

	output := []string{checks, httpReqFailed, dataReceived, dataSent, avg, min, med, max, p90, p95}
	fmt.Println(strings.Join(output, "\t"))
}

func extractPercentage(line string) string {
	re := regexp.MustCompile(`(\d+\.\d+%)`)
	match := re.FindStringSubmatch(line)
	if len(match) > 0 {
		return match[0]
	}
	return ""
}

func extractData(line string) string {
	re := regexp.MustCompile(`:\s+([\d.]+\s\w+)`)
	match := re.FindStringSubmatch(line)
	if len(match) > 0 {
		return match[1]
	}
	return ""
}

func extractValue(line, key string) string {
	re := regexp.MustCompile(key + `=([\d.]+\w*)`)
	match := re.FindStringSubmatch(line)
	if len(match) > 0 {
		return match[1]
	}
	return ""
}
