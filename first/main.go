package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ns = map[string]string{
	"nol":      "0",
	"satu":     "1",
	"dua":      "2",
	"tiga":     "3",
	"empat":    "4",
	"lima":     "5",
	"enam":     "6",
	"tujuh":    "7",
	"delapan":  "8",
	"sembilan": "9",
	"sepuluh":  "10",
	"sebelas":  "11",
	"seratus":  "100",
	"puluh":    "0",
	"ratus":    "00",
}

var sn = map[string]string{
	"0": "nol",
	"1": "satu",
	"2": "dua",
	"3": "tiga",
	"4": "empat",
	"5": "lima",
	"6": "enam",
	"7": "tujuh",
	"8": "delapan",
	"9": "sembilan",
}

func parseSentence(text string) string {
	split := strings.Split(text, " ")

	var x []string
	var op []string
	var v string

	for i := 0; i < len(split); i++ {
		switch split[i] {
		case "tambah":
			x = append(x, v)
			op = append(op, "+")
			v = ""
		case "kurang":
			x = append(x, v)
			op = append(op, "-")
			v = ""
		case "kali":
			x = append(x, v)
			op = append(op, "x")
			v = ""
		case "bagi":
			x = append(x, v)
			op = append(op, "/")
			v = ""
		case "belas":
			v = "1" + v
		case "puluh":
			if i+1 < len(split) {
				if ns[split[i+1]] == "" {
					v = v + "0"
				}
			} else {
				v = v + "0"
			}
		default:
			v = v + ns[split[i]]
		}

		if i == len(split)-1 {
			x = append(x, v)
		}
	}

	v = doMathOperation(x, op)

	return getNumberWord(v)
}

func getNumberWord(a string) string {
	ans := ""

	if string(a[0]) == "-" {
		ans = "negatif "
		a = a[1:len(a)]
	}

	num, _ := strconv.ParseInt(a, 10, 0)

	switch {
	case num > 1000:
		if a[1:len(a)] == "000" {
			ans = ans + "seribu"
		} else {
			// Handle three digits
		}
	case num > 199:
		if a[1:len(a)] == "00" {
			ans = ans + sn[string(a[0])] + " ratus"
		} else {
			// Handle two digits
		}
	case num > 100:
		if a[1:len(a)] == "00" {
			ans = ans + "seratus"
		} else {
			// Handle two digits
		}
	case num > 19:
		if string(a[1]) == "0" {
			ans = ans + sn[string(a[0])] + " puluh"
		} else {
			ans = ans + sn[string(a[0])] + " puluh " + sn[string(a[1])]
		}
	case num > 11:
		ans = ans + sn[string(a[1])] + " belas"
	case num == 11:
		ans = ans + "sebelas"
	case num == 10:
		ans = ans + "sepuluh"
	default:
		ans = ans + sn[string(a[0])]
	}

	return ans
}

func doMathOperation(a []string, b []string) string {
	ans, _ := strconv.ParseInt(a[0], 10, 0)

	for i := 0; i < len(b); i++ {
		switch b[i] {
		case "+":
			x, _ := strconv.ParseInt(a[i+1], 10, 0)
			ans += x
		case "-":
			x, _ := strconv.ParseInt(a[i+1], 10, 0)
			ans -= x
		case "x":
			x, _ := strconv.ParseInt(a[i+1], 10, 0)
			ans *= x
		case "/":
			x, _ := strconv.ParseInt(a[i+1], 10, 0)
			ans /= x
		}
	}

	return strconv.FormatInt(ans, 10)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	ans := parseSentence(text)

	fmt.Println("Output: " + text + " adalah " + ans)
}
