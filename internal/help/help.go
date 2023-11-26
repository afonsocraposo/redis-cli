package help

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"github.com/afonsocraposo/redis-cli/configs"
	"github.com/afonsocraposo/redis-cli/internal/helpers"
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

func getBlockHelp(arg argument) string {
	options := make([]string, len(*arg.Arguments))
	for i, option := range *arg.Arguments {
		h := parseHelpArg(option)
		options[i] = h
	}
	h := strings.Join(options, " ")
	if arg.Optional != nil && *arg.Optional && arg.Token != nil {
		return fmt.Sprintf("[%s %s]", *arg.Token, h)
	} else {
		return fmt.Sprintf("[%s]", h)
	}
}

func getOneOfHelp(arg argument) string {
	options := make([]string, len(*arg.Arguments))
	for i, option := range *arg.Arguments {
		h := fmt.Sprintf("%s", *option.Token)
		if option.Type != "pure-token" {
			h += fmt.Sprintf(" %s", getPureToken(option))
		}
		options[i] = h
	}
	help := fmt.Sprintf("[%s]", strings.Join(options, "|"))
	return help
}

func getPureToken(arg argument) string {
	if arg.Optional != nil && *arg.Optional && arg.Token != nil {
		return fmt.Sprintf("[%s]", *arg.Token)
	} else {
		return fmt.Sprintf("%s", arg.Name)
	}
}

func getDefault(arg argument) string {
	if arg.Optional != nil && *arg.Optional && arg.Token != nil {
		return fmt.Sprintf("[%s %s]", *arg.Token, arg.Name)
	} else {
		return fmt.Sprintf("%s", arg.Name)
	}
}

func parseHelpArg(arg argument) string {
	switch arg.Type {
	case "oneof":
		return getOneOfHelp(arg)
	case "block":
		return getBlockHelp(arg)
	case "pure-token":
		return getPureToken(arg)
	default:
		return getDefault(arg)
	}
}

func GetHelpText(command string) string {
    helpPath := helpers.ParsePath((settings.HELP_ASSETS_PATH))

	lCommand := strings.ToLower(command)
	filepath := path.Join(helpPath, lCommand) + ".json"
	c, err := parseFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	uCommand := strings.ToUpper(command)
	help := "\n" + uCommand
	for _, arg := range c.Arguments {
		h := parseHelpArg(arg)
		help += fmt.Sprintf(" %s", h)
	}
	help += fmt.Sprintf("\nsummary: %s\nsince: %s\ngroup: %s\n", c.Summary, c.Since, c.Group)
	return help
}
