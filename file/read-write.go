package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readCSV()
	writeText()
	createFile()
	createFileAndInsertData()
}

// NOTE - read csv
func readCSV() {
	file, err := os.Open("file/students.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// }

	// count := 1
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Printf("Line %d : %s\n", count, line)
	// 	count++
	// }

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")
		fmt.Println(item[0])
		count++
	}
}

// NOTE - insert data into txt file
func writeText() {
	data1 := []byte(" Hello\n My name is Ton\n I am 22 Year old")
	err := os.WriteFile("file/data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}
}

func createFile() {
	f, ferr := os.Create("file/StudentName")
	if ferr != nil {
		panic(ferr)
	}

	defer f.Close()
}

func createFileAndInsertData() {
	data2 := []byte(" Hello\n My name is Phon\n I am 24 Year old")
	os.WriteFile("file/StudentName.txt", data2, 0644)
}
