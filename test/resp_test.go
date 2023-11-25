package test

import (
	"github.com/afonsocraposo/redis-cli/internal/resp"
	"testing"
)

type test struct {
	input, expected string
}

var deserialiseTests = []test{
	{"$-1\r\n", "null"},
	{"*-1\r\n", "null"},
	{"*1\r\n$4\r\nping\r\n", "[ping]"},
	{"*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n", "[echo, hello world]"},
	{"*2\r\n$4\r\necho\r\n$-1\r\n", "[echo, null]"},
	{"*2\r\n$3\r\nget\r\n$3\r\nkey\r\n", "[get, key]"},
	{"*3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n", "[set, key, value]"},
	{"+OK\r\n", "OK"},
	{"-Error message\r\n", "Error: Error message"},
	{"$0\r\n\r\n", ""},
	{"+hello world\r\n", "hello world"},
	{":0\r\n", "0"},
	{":1000\r\n", "1000"},
	{":-1000\r\n", "-1000"},
	{"_\r\n", "null"},
}

var serialiseTests = []test{
	{"ping", "*1\r\n$4\r\nping\r\n"},
	{"echo \"hello world\"", "*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n"},
	{"get key", "*2\r\n$3\r\nget\r\n$3\r\nkey\r\n"},
	{"set key value", "*3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"},
}

func TestDeserialise(t *testing.T) {
	for _, test := range deserialiseTests {
		if output := resp.Deserialise(test.input); output != test.expected {
			t.Errorf("Output %s not equal to expected %s", output, test.expected)
		}
	}
}

func TestSerialise(t *testing.T) {
	for _, test := range serialiseTests {
		if output := resp.Serialise(test.input); output != test.expected {
			t.Errorf("Output %s not equal to expected %s", output, test.expected)
		}
	}
}
