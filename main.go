package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yalp/jsonpath"
	"os"
	"runtime"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"

// var Cyan = "\033[36m"
var Gray = "\033[37m"

//var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		//Cyan = ""
		Gray = ""
		//White = ""
	}
}

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
		fmt.Println(string(line))
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
		if err != nil || js == "" {
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
	s = strings.TrimLeft(s, " ")

	// color
	if config.Color {
		level, _ := jsonpath.Read(logline, "$.level")
		s = colorize(fmt.Sprintf("%s", level), s)
	}

	fmt.Println(s)
}

func colorize(level string, line string) string {
	switch strings.ToUpper(level) {
	case "INFO":
		return fmt.Sprintf("%s%s%s", Green, line, Reset)
	case "WARN":
		return fmt.Sprintf("%s%s%s", Yellow, line, Reset)
	case "ERROR":
		return fmt.Sprintf("%s%s%s", Red, line, Reset)
	case "DEBUG":
		return fmt.Sprintf("%s%s%s", Blue, line, Reset)
	case "TRACE":
		return fmt.Sprintf("%s%s%s", Purple, line, Reset)
	default:
		return fmt.Sprintf("%s%s%s", Gray, line, Reset)
	}
}

func exit(err error, code int) {
	fmt.Println(err)
	os.Exit(code)
}
