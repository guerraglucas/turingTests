// There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.
// Example
// 		n = 7
// 		ar = [1,2,1,2,1,3,2]
// There is one pair of color  and one of color . There are three odd socks left, one of each color. The number of pairs is .
// Function Description
// Complete the sockMerchant function in the editor below.
// sockMerchant has the following parameter(s):
// 		int n: the number of socks in the pile
// 		int ar[n]: the colors of each sock
// Returns
// 		int: the number of pairs
// Input Format
// 		The first line contains an integer , the number of socks represented in .
// 		The second line contains  space-separated integers, , the colors of the socks in the pile.

package solution

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here

	mapOfSocks := make(map[int32]int32)
	var amountOfPairs int32 = 0

	for _, color := range ar {
		mapOfSocks[color]++
	}

	for _, amount := range mapOfSocks {
		numberOfPairsOfThisColor := amount / 2
		amountOfPairs += numberOfPairsOfThisColor
	}

	return amountOfPairs
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

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
