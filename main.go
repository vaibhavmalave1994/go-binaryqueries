package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var i int64
	fmt.Println("Binary Queries problem started")
	inputScan := bufio.NewScanner(os.Stdin)
	outputScan := bufio.NewWriter(os.Stdout)
	fmt.Println("Enter length of binary array(N)")
	inputScan.Split(bufio.ScanWords)
	inputScan.Scan()
	n, err := strconv.ParseInt(inputScan.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter Number of queries")
	inputScan.Scan()
	q, err := strconv.ParseInt(inputScan.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	arr := make([]int64, n)

	fmt.Printf("\nPlease enter array with size (%d):\n", n)
	for i = 0; i < n; i++ {
		inputScan.Scan()
		c, _ := strconv.ParseInt(inputScan.Text(), 10, 64)
		arr[i] = c
	}
	fmt.Println("write query")
	for q > 0 {
		inputScan.Scan()
		c, _ := strconv.ParseInt(inputScan.Text(), 10, 64)
		if c == 0 {
			inputScan.Scan()
			inputScan.Scan()
			index, _ := strconv.ParseInt(inputScan.Text(), 10, 64)
			if arr[index-1] == 0 {
				outputScan.WriteString("EVEN\n")
			} else {
				outputScan.WriteString("ODD\n")
			}
		} else {
			inputScan.Scan()
			index, _ := strconv.ParseInt(inputScan.Text(), 10, 64)
			if arr[index-1] == 0 {
				arr[index-1] = 1
			} else {
				arr[index-1] = 0
			}
		}
		q--
	}
	outputScan.Flush()
}
