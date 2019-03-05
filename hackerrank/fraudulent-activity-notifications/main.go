package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the activityNotifications function below.

func activityNotifications(expenditure []int32, d int32) int32 {
	notificationCounter := int32(0)
	hasEasyMedian := len(expenditure)%2 == 1

	if hasEasyMedian {
		medianIndex := d / 2
		for i := int32(0); i+d < int32(len(expenditure)); i++ {
			if expenditure[i+d] >= expenditure[medianIndex+i] {
				notificationCounter++
			}
		}
	} else {
		for i := int32(0); i+d < int32(len(expenditure)); i++ {
			if float64(expenditure[i+d]) >= float64(expenditure[(d/2)+i]+expenditure[(d/2)+1+i]) {
				notificationCounter++
			}
		}
	}
	return notificationCounter
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
