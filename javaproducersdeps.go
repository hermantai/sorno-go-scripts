package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func printUsage() {
	fmt.Printf("Usage: %s input_file\n", filepath.Base(os.Args[0]))
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	fmt.Println("Process", os.Args[1])

	fileContent, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}

	remaining := string(fileContent)
	producerFunc := ""
	for {
		producerFunc, remaining = getProducerFunc(remaining)
		if producerFunc == "" {
			break
		}
		fmt.Printf("producerFun: %s\n", producerFunc)
	}

	os.Exit(0)
}

func getProducerFunc(s string) (producerFunc string, remaining string) {
	//r := regexp.MustCompile(`(?P<leadingSpaces> *)@Produces.*`)
	params := getParams(`(?s)(?P<leadingSpaces> *)@Produces.*`, s)
	if params == nil {
		return "", ""
	}

	r := regexp.MustCompile(fmt.Sprintf(".*\n%s}", params["leadingSpaces"]))
	matchedString := params[""]

	loc := r.FindStringIndex(matchedString)
	if loc == nil {
		return "", ""
	}

	return matchedString[0:loc[1]], matchedString[loc[1]:len(matchedString)]
}

func getParams(regEx, s string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(s)

	if match == nil {
		return nil
	}

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
