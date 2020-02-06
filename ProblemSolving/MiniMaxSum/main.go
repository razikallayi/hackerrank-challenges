package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.

func miniMaxSum(arr []int32) {
	var len = len(arr)
	var sum []int64 = make([]int64, len)
	var min int64 = 0
	var max int64 = 0
	var sub []int32 = make([]int32, 0, 4)
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if i != j {
				sub = append(sub, arr[j])
			}
		}
		sum[i] = arraySum(sub)
		sub = make([]int32, 0, 4)
	}

	for s := range sum {
		if s == 0 || sum[s] < min {
			min = sum[s]
		}
		if s == 0 || sum[s] > max {
			max = sum[s]
		}
	}
	fmt.Printf("%d %d", min, max)
}

func arraySum(arr []int32) int64 {
	var sum int64 = 0
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
	}
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
