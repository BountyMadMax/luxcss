package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

func Extract(directories []string, fileExtensions []string, prefix string) {
	fmt.Println("--- Start extracting ---")
	fmt.Println("Prefix: ", prefix)
	fmt.Println("Directories: ", directories)
	fmt.Println("File extensions: ", fileExtensions)

	breakpoints := Breakpoints()
	breakpointNames := make([]string, len(breakpoints))
	for i, breakpoint := range breakpoints {
		breakpointNames[i] = breakpoint.name
	}

	states := States()
	stateNames := make([]string, len(states))
	for i, state := range states {
		stateNames[i] = state.name
	}

	styles := Styles()
	stylePatterns := make([]string, len(styles))
	for i, style := range styles {
		stylePatterns[i] = style.name
	}

	fileRegex, e := regexp.Compile(fmt.Sprintf(".(%s)$", strings.Join(fileExtensions, "|")))
	if e != nil {
		log.Fatal(e)
	}

	styleRegex, e := regexp.Compile(fmt.Sprintf("((%s):)*((%s):)*(%s)", strings.Join(breakpointNames, "|"), strings.Join(stateNames, "|"), strings.Join(stylePatterns, "|")))

	fmt.Println("File regex: ", fileRegex)
	fmt.Println("Style regex: ", styleRegex)

	fmt.Println("---------")
	for _, dir := range directories {
		extractDir(dir, fileRegex, styleRegex)
	}
}

func extractDir(dir string, fileRegex *regexp.Regexp, styleRegex *regexp.Regexp) {
	e := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil && fileRegex.MatchString(info.Name()) {
			println(info.Name())

			extractFile(path, styleRegex)
		}
		return nil
	})

	if e != nil {
		log.Fatal(e)
	}
}

func extractFile(filePath string, regex *regexp.Regexp) []string {
	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	regexResult := regex.FindAllString(string(content[:]), 100000)
	result := make([]string, 0)

	for _, res := range regexResult {
		if !slices.Contains(result, res) {
			result = append(result, res)
			fmt.Println("Added:", res)
		}
	}

	fmt.Println("Uniques:", result)

	return result
}
