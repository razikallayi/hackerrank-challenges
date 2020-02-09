package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	var count int32 = -1
	len := len(c)
	for i := 0; i < len; i, count = i+1, count+1 {
		if i < len-2 && c[i+2] == 0 {
			i++
		}
	}
	return count
}

func main() {
	result := jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0})
	fmt.Printf("%+v\n", result)
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	// nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// cTemp := strings.Split(readLine(reader), " ")

	// var c []int32

	// for i := 0; i < int(n); i++ {
	// 	cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
	// 	checkError(err)
	// 	cItem := int32(cItemTemp)
	// 	c = append(c, cItem)
	// }

	// result := jumpingOnClouds(c)

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
