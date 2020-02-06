package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
	var max int32 = 0
	var count int32 = 0
	var i int32 = 0
	for ; i < int32(len(ar)); i++ {
		if max != 0 || ar[i] > max {
			max = ar[i]
		}
	}
	fmt.Println(max)
	i = 0
	for ; i < int32(len(ar)); i++ {
		if ar[i] == max {
			count++
		}
	}
	fmt.Println(count)
	return count
}

func main() {
	birthdayCakeCandles([]int32{3, 2, 1, 3})
	// reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	// arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)

	// arTemp := strings.Split(readLine(reader), " ")

	// var ar []int32

	// for i := 0; i < int(arCount); i++ {
	//     arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
	//     checkError(err)
	//     arItem := int32(arItemTemp)
	//     ar = append(ar, arItem)
	// }

	// result := birthdayCakeCandles(ar)

	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
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
