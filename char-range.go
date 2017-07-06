package main

import (
	"time"
	"fmt"
	"strings"
	"errors"
	"flag"
)

type ReversableString string

func (str ReversableString) StringReverse() ReversableString {
	rev := ReversableString("")
	for i := len(str) - 1; i >= 0; i-- {
		rev += ReversableString(str[i])
	}
	return rev
}

func ComparaSecuencias(a string, b string) int {
	var retorno int
	if len(a) != len(b) {
		retorno = len(a) - len(b)
	} else {
		switch {
			case a < b:
				retorno = -1
				break
			case a > b:
				retorno = 1
				break
			default:
				retorno = 0
				break
		}
	}
	return retorno
}

func NextChar(ch string) (int, string, error) {
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var index int
	var character string
	for i, c := range chars {
		if c == strings.ToLower(ch) {
			if len(chars)-1 == i {
				index = 0
			} else {
				index = i + 1
			}
			character = chars[index]
			break
		}
	}

	if index == 0 && character == "" {
		return index, character, errors.New("The requested character is not in the list.")
	}

	return index, character, nil

}

func NextSequence(chsq ReversableString) ReversableString {
	rev := chsq.StringReverse()
	newsq := ReversableString("")
	for j := 0; j <= len(rev)-1; j++ {
		ni, nc, _ := NextChar(string(rev[j]))
		newsq += ReversableString(nc)
		if ni != 0 {
			newsq += rev[j+1:]
			break
		}
		if ni == 0 && j == len(chsq)-1 {
			newsq += ReversableString(nc)
		}
	}
	return newsq.StringReverse()
}

func main() {

	startStr := flag.String("start", "a", "Start string")
	endStr := flag.String("end", "z", "End string")

	flag.Parse()

	if comp := ComparaSecuencias(*startStr,*endStr); comp >= 0 {
		panic("A cadea de fin debe ser maior que a cadea de inicio.")
	}

	generator := make(chan ReversableString)

	go func(start ReversableString, end ReversableString) {
		generator <- start
		qc := start
		for {
			nextsq := NextSequence(qc)
			generator <- nextsq
			if nextsq != end {
				qc = nextsq
			} else {
				close(generator)
				break
			}
		}
	}(ReversableString(*startStr),ReversableString(*endStr))

	for nextsq := range generator {
		fmt.Println(string(nextsq))
		time.Sleep(25 * time.Millisecond)
	}
}