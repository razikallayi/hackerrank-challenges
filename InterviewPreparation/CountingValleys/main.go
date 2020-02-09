package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var level int32 = 0
	var belowSeaLevel bool = false
	var valleyCount int32 = 0
	for i := range s {
		if s[i] == 'U' {
			level++
		} else {
			level--
		}

		if level < 0 && belowSeaLevel == false {
			belowSeaLevel = true
		}

		if level == 0 && belowSeaLevel == true {
			valleyCount++
			belowSeaLevel = false
		}
	}
	return valleyCount
}
func main() {
	var n int32 = 8
	var s string = "DDUUDDUDUUUD"
	// var s string = "DDUUUUDD"

	//              /\
	//    \  /\    /
	// 	   \/  \/\/

	result := countingValleys(n, s)
	fmt.Printf("%+v\n", result)
	// reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	// nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// s := readLine(reader)

	// result := countingValleys(n, s)

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
