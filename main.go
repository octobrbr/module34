package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var res int

	if len(os.Args) != 3 {
		fmt.Println("Не введены параметры")
		return
	}

	fileIn := os.Args[1]
	fileOut := os.Args[2]

	_ = os.Remove(fileOut)
	fOut, err := os.OpenFile(fileOut, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer fOut.Close()
	writer := bufio.NewWriter(fOut)

	fIn, err := os.ReadFile(fileIn)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`([0-9]+)([\+\-\\\*])+([0-9]+)(\=)`)

	submatches := re.FindAllStringSubmatch(string(fIn), -1)
	
	writer.Flush()
	for _, v := range submatches {
		switch v[2] {
		case "+":
			a, err := strconv.Atoi(v[1])
			if err != nil {
				panic(err)
			}

			b, err := strconv.Atoi(v[3])
			if err != nil {
				panic(err)
			}

			res = a + b
		case "-":
			a, err := strconv.Atoi(v[1])
			if err != nil {
				panic(err)
			}

			b, err := strconv.Atoi(v[3])
			if err != nil {
				panic(err)
			}

			res = a - b
		}

		writer.WriteString(v[0] + fmt.Sprint(res) + "\n")
	}
	writer.Flush()
}
