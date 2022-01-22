package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {

	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(file)

	var index int = 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		numbers[index], err = strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return numbers, err
		}

		index++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, nil
}