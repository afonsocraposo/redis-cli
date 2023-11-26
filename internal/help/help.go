package help

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type command struct {
	Summary         string      `json:"summary"`
	Complexity      string      `json:"complexity"`
	Group           string      `json:"group"`
	Since           string      `json:"since"`
	Arity           int         `json:"arity"`
	Function        string      `json:"function"`
	GetKeysFunction string      `json:"get_keys_function"`
	History         [][]string  `json:"history"`
	CommandFlags    []string    `json:"command_flags"`
	AclCategories   []string    `json:"acl_categories"`
	KeySpecs        []keySpec   `json:"key_specs"`
	ReplySchema     replySchema `json:"reply_schema"`
	Arguments       []argument  `json:"arguments"`
}

type keySpec struct {
	Notes       string      `json:"notes"`
	Flags       []string    `json:"flags"`
	BeginSearch beginSearch `json:"begin_search"`
	FindKeys    findKeys    `json:"find_keys"`
}

type index struct {
	Pos int `json:"pos"`
}

type myRange struct {
	LastKey int `json:"lastkey"`
	Step    int `json:"step"`
	Limit   int `json:"limit"`
}

type beginSearch struct {
	Index index `json:"index"`
}

type findKeys struct {
	Range myRange `json:"range"`
}

type replySchema struct {
	AnyOf []anyOf `json:"anyOf"`
}

type anyOf struct {
	Description string  `json:"description"`
	Type        *string `json:"type,omitempty"`
	Const       *string `json:"const,omitempty"`
}

type argument struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	KeySpecIndex *int        `json:"key_sec_index,omitempty"`
	Optional     *bool       `json:"optional,omitempty"`
	Since        *string     `json:"since,omitempty"`
	Token        *string     `json:"token,omitempty"`
	Arguments    *[]argument `json:"arguments,omitempty"`
}

func getFilename(filepath string) string {
	filenameWithExt := path.Base(filepath)
	ext := path.Ext(filenameWithExt)
	filename := strings.TrimSuffix(filenameWithExt, ext)
	return filename
}

func parseFile(filepath string) (*command, error) {
	filename := getFilename(filepath)
	upperFilename := strings.ToUpper(filename)

	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var command map[string]command

	err = json.Unmarshal(byteValue, &command)
	if err != nil {
		log.Fatal(err)
	}

	c := command[upperFilename]
	return &c, nil
}

func GetHelpText(command string, helpPath string) string {
	lCommand := strings.ToLower(command)
    filepath := path.Join(helpPath, lCommand) + ".json"
	c, err := parseFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	uCommand := strings.ToUpper(command)
	help := "\n"+uCommand
	for _, arg := range c.Arguments {
		switch arg.Type {
		case "oneof":
			options := make([]string, len(*arg.Arguments))
			for i, option := range *arg.Arguments {
				h := fmt.Sprintf("%s", *option.Token)
				if option.Type != "pure-token" {
					h += fmt.Sprintf(" %s", option.Name)
				}
				options[i] = h
			}
			help += fmt.Sprintf(" [%s]", strings.Join(options, "|"))
		default:
			help += fmt.Sprintf(" %s", arg.Name)
		}
	}
    help += fmt.Sprintf("\nsummary: %s\nsince: %s\ngroup: %s\n", c.Summary, c.Since, c.Group)
	return help
}

