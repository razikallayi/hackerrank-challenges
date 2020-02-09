package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var applefall []int32
	var orangefall []int32
	var appleCount int32
	var orangeCount int32
	for i := range apples {
		applefall = append(applefall, apples[i]+a)
	}
	for i := range oranges {
		orangefall = append(orangefall, oranges[i]+b)
	}
	for i := range applefall {
		if applefall[i] >= s && applefall[i] <= t {
			appleCount++
		}
	}
	for i := range orangefall {
		if orangefall[i] >= s && orangefall[i] <= t {
			orangeCount++
		}
	}
	fmt.Println(appleCount)
	fmt.Println(orangeCount)
}

func main() {
	var s int32 = 7
	var t int32 = 11
	var a int32 = 5
	var b int32 = 15
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}
	countApplesAndOranges(s, t, a, b, apples, oranges)
	// reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// st := strings.Split(readLine(reader), " ")

	// sTemp, err := strconv.ParseInt(st[0], 10, 64)
	// checkError(err)
	// s := int32(sTemp)

	// tTemp, err := strconv.ParseInt(st[1], 10, 64)
	// checkError(err)
	// t := int32(tTemp)

	// ab := strings.Split(readLine(reader), " ")

	// aTemp, err := strconv.ParseInt(ab[0], 10, 64)
	// checkError(err)
	// a := int32(aTemp)

	// bTemp, err := strconv.ParseInt(ab[1], 10, 64)
	// checkError(err)
	// b := int32(bTemp)

	// mn := strings.Split(readLine(reader), " ")

	// mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	// checkError(err)
	// m := int32(mTemp)

	// nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// applesTemp := strings.Split(readLine(reader), " ")

	// var apples []int32

	// for i := 0; i < int(m); i++ {
	//     applesItemTemp, err := strconv.ParseInt(applesTemp[i], 10, 64)
	//     checkError(err)
	//     applesItem := int32(applesItemTemp)
	//     apples = append(apples, applesItem)
	// }

	// orangesTemp := strings.Split(readLine(reader), " ")

	// var oranges []int32

	// for i := 0; i < int(n); i++ {
	//     orangesItemTemp, err := strconv.ParseInt(orangesTemp[i], 10, 64)
	//     checkError(err)
	//     orangesItem := int32(orangesItemTemp)
	//     oranges = append(oranges, orangesItem)
	// }

	// countApplesAndOranges(s, t, a, b, apples, oranges)
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
