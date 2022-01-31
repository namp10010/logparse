package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yalp/jsonpath"
	"os"
	"strings"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		exit(err, 1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		printLine(line, config)
	}

}

func printLine(line []byte, config *Config) {
	var logline interface{}
	if err := json.Unmarshal(line, &logline); err != nil {
		fmt.Printf("logparse: %s\n", err)
		return
	}

	// skip
	// can the performance be improved
	for _, skippedText := range config.Skips {
		// could do regex
		if strings.Contains(string(line), skippedText) {
			return
		}
	}

	// format
	var s string
	for _, f := range config.Fields {
		js, err := jsonpath.Read(logline, f)
		if err != nil {
			continue
		}
		s = fmt.Sprintf("%s %s", s, js)
	}
	for _, f := range config.JsonFields {
		js, err := jsonpath.Read(logline, f)
		if err != nil {
			continue
		}

		b, err := json.Marshal(js)
		if err == nil {
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, b, "", "  "); err == nil {
				js = prettyJSON.String()
			}
		}

		s = fmt.Sprintf("%s %s", s, js)
	}
	fmt.Println(s)
}

func exit(err error, code int) {
	fmt.Println(err)
	os.Exit(code)
}
