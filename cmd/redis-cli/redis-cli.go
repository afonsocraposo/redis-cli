package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/afonsocraposo/redis-cli/internal/help"
	"github.com/afonsocraposo/redis-cli/internal/resp"
	"github.com/afonsocraposo/redis-cli/internal/tcp"
	"log"
	"os"
	"strings"
)

const HELP_PATH = "/Users/prezi/Documents/redis-cli/assets/help/commands/"

func main() {
	var hostname string
	var port int

	flag.StringVar(&hostname, "h", "localhost", "hostname")
	flag.IntVar(&port, "p", 6379, "port")
	flag.Parse()

	address := fmt.Sprintf("%s:%d", hostname, port)

	fmt.Println("Go Redis CLI")

	client := &tcp.TCPClient{}
	err := client.Connect(address)
	if err != nil {
		log.Fatalf("Error connecting to %s\n%v", address, err)
	}

	// disconnect from server on exit
	defer func() {
		err := client.Disconnect()
		if err != nil {
			log.Printf("Disconnection error: %v", err)
		}
	}()

	// user input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s> ", address)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("An error occurred:", err)
		}
		// Trim the newline character from the input
		input = strings.TrimSpace(input)

		if input == "quit" {
			return
		}

		if strings.HasPrefix(input, "help") {
			after, _ := strings.CutPrefix(input, "help")
			command := strings.TrimSpace(after)
			if after == "" {
				fmt.Println("HELP COMMAND")
			} else {
				helpText := help.GetHelpText(command, HELP_PATH)
				fmt.Println(helpText)
			}
		} else {

			// proccess input
			r := resp.Serialise(input)
			client.Write(r)

			reply, err := client.Read()
			if err != nil {
				log.Fatal("Error reading reply:", err)
			}
			fmt.Println(resp.Deserialise(reply))
		}
	}
}
