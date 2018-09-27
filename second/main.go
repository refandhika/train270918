package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func highestSum(a string) string {
	rep := strings.NewReplacer("[", "", "]", "")
	a = rep.Replace(a)

	spl := strings.Split(a, " ")
	arr := []int{}

	for _, v := range spl {
		num, _ := strconv.ParseInt(v, 10, 0)
		arr = append(arr, int(num))
	}

	max := 0
	for i := 0; i < len(arr)-1; i++ {
		for j, k := 0, i; j < i+1; j++ {
			sum := 0
			for _, v := range arr[j : len(arr)-k] {
				sum += v
			}
			if max < sum {
				max = sum
			}
			k--
		}
	}

	return strconv.Itoa(max)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	ans := highestSum(text)

	fmt.Println("Output: " + ans)
}
