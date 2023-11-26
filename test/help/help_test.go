package help

import (
	"github.com/afonsocraposo/redis-cli/internal/help"
	"testing"
)

type test struct {
	input, expected string
}

const setHelp = "\nSET key value [NX|XX] [GET] [EX seconds|PX milliseconds|EXAT unix-time-seconds|PXAT unix-time-milliseconds|KEEPTTL]\nsummary: Sets the string value of a key, ignoring its type. The key is created if it doesn't exist.\nsince: 1.0.0\ngroup: string\n"
const helloHelp = "\nHELLO [protover [AUTH username password] [SETNAME clientname]]\nsummary: Handshakes with the Redis server.\nsince: 6.0.0\ngroup: connection\n"

var getHelpTextTests = []test{
	{"set", setHelp},
	{"hello", helloHelp},
}

func TestGetHelpText(t *testing.T) {
	for _, test := range getHelpTextTests {
		if output := help.GetHelpText(test.input); output != test.expected {
			t.Errorf("Output %s not equal to expected %s", output, test.expected)
		}
	}
}
