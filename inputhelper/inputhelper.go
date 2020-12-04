package inputhelper

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetStrings(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var outputs []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		outputs = append(outputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return outputs
}

func Unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

var delim = []byte{'\n', '\n'}

func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	for i := 0; i+len(delim) <= len(data); {
		j := i + bytes.IndexByte(data[i:], delim[0])
		if j < i {
			break
		}
		if bytes.Equal(data[j+1:j+len(delim)], delim[1:]) {
			// We have a full delim-terminated line.
			return j + len(delim), data[0:j], nil
		}
		i = j + 1
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func GetFileAsString(x string) string {
	content, err := ioutil.ReadFile(x)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	return string(content)
}

func GetStringsFromFileByDelim(x string) []string {
	str := GetFileAsString(x)
	var out []string

	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(ScanLines)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	return out
}
