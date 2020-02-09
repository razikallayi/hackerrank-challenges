package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	chaos := false
	var count int32 = 0
	var len = len(q)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len; i++ {
			if q[i-1]-int32(i) > 2 {
				chaos = true
				break
			}
			if q[i-1] > q[i] {
				q[i-1], q[i] = q[i], q[i-1]
				swapped = true
				count++
			}
		}
	}
	if chaos {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(count)
	}
}

func elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("%v\n", time.Since(start))
	}
}

func main() {
	defer elapsed()()
	time.Sleep(time.Second * 2)
	minimumBribes([]int32{2, 1, 5, 3, 4})
	minimumBribes([]int32{2, 5, 1, 3, 4})
	minimumBribes(test8)

	// reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	// tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)
	// t := int32(tTemp)

	// for tItr := 0; tItr < int(t); tItr++ {
	//     nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	//     checkError(err)
	//     n := int32(nTemp)

	//     qTemp := strings.Split(readLine(reader), " ")

	//     var q []int32

	//     for i := 0; i < int(n); i++ {
	//         qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
	//         checkError(err)
	//         qItem := int32(qItemTemp)
	//         q = append(q, qItem)
	//     }

	//     minimumBribes(q)
	// }
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
