package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the designerPdfViewer function below.
func designerPdfViewer(h []int32, word string) int32 {
	var subset []int32
	for letter := range word {
		position := int32(word[letter]) - 97 // ascii value
		subset = append(subset, h[position])
	}
	return int32(len(word)) * max(subset)
}

func max(arr []int32) int32 {
	var max int32 = 0
	for i := range arr {
		if max == 0 || arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	word := "abc"
	result := designerPdfViewer(h, word)
	fmt.Printf("%+v\n", result)
	// reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	// hTemp := strings.Split(readLine(reader), " ")

	// var h []int32

	// for i := 0; i < 26; i++ {
	//     hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
	//     checkError(err)
	//     hItem := int32(hItemTemp)
	//     h = append(h, hItem)
	// }

	// word := readLine(reader)

	// result := designerPdfViewer(h, word)

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
