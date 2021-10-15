package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	var m map[string]interface{}
	if err := json.Unmarshal(line, &m); err != nil {
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

	// print
	var s string
	for _, f := range config.Fields {
		if v, ok := m[f]; !ok {
			continue
		} else {
			// can improve with width
			js,err := json.Marshal(v)
			if err != nil {
				fmt.Printf("logparse: %s\n", err)
				continue
			}
			s = fmt.Sprintf("%s %s", s, js)
		}
	}
	fmt.Println(s)
}

func exit(err error, code int) {
	fmt.Println(err)
	os.Exit(code)
}
