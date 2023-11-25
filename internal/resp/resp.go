package resp

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Deserialise(input string) string {
	after := input
	output := ""
	for after != "" && strings.Contains(after, "\r\n") {
		var o string
		o, after = deserialise(after)
		output += o
	}
	return output
}

func deserialise(input string) (string, string) {
	if input == "" {
		return "", ""
	}
	output := ""
	after := ""
	switch input[0] {
	case '-':
		output += "Error: "
		fallthrough
	case '+':
		before, a, _ := strings.Cut(input[1:], "\r\n")
		after = a
		output += fmt.Sprintf("%s", before)
	case ':':
		before, a, _ := strings.Cut(input[1:], "\r\n")
		after = a
		output = before
	case '$':
		before, a, _ := strings.Cut(input[1:], "\r\n")
		after = a
		n, err := strconv.Atoi(before)
		if err != nil {
			log.Println(err)
			break
		}
		if n == -1 {
			output += "null"
			break
		}
		output += fmt.Sprintf("%s", a[:n])
		_, a, _ = strings.Cut(a[n:], "\r\n")
		after = a
	case '*':
		before, a, _ := strings.Cut(input[1:], "\r\n")
		n, err := strconv.Atoi(before)
		if err != nil {
			log.Println(err)
			break
		}
		if n == -1 {
			output += "null"
			break
		}
		parts := make([]string, n)
		aa := a
		for i := 0; i < n; i++ {
			o, _aa := deserialise(aa)
			aa = _aa
			parts[i] = o
		}
		after = aa
		output = fmt.Sprintf("[%s]", strings.Join(parts, ", "))
	case '_':
		output = "null"
		_, a, _ := strings.Cut(input[1:], "\r\n")
		after = a
	}
	return output, after
}

func Serialise(input string) string {
	after := input
	output := ""
    c := 0
	for after != "" {
		var o string
		o, after = serialise(after)
		output += o
        c++
	}
	return fmt.Sprintf("*%d\r\n%s", c, output)
}

func serialise(input string) (string, string) {
	if input == "" {
		return "", ""
	}
    var before, after string
    if input[0] == '"' {
        before, after, _ = strings.Cut(input[1:], "\"")
    } else {
        before, after, _ = strings.Cut(input, " ")
    }
    r := fmt.Sprintf("$%d\r\n%s\r\n", len(before), before)
    return r, after
}
