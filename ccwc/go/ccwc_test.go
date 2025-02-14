package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestWordCounter(t *testing.T) {
	getFilePath := func(fileName string) string {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		testFilePath, err := filepath.Abs(filepath.Join(cwd, "../"+fileName))

		if err != nil {
			panic(err)
		}
		return testFilePath
	}

	countWithWC := func(flag, testFilePath string) int64 {
		out, err := exec.Command("wc", flag, testFilePath).Output()
		if err != nil {
			panic(err)
		}

		byteCounted := strings.Split(string(out), " ")[0]
		expected, err := strconv.ParseInt(byteCounted, 10, 64)
		if err != nil {
			panic(err)
		}
		return expected
	}

	type TestStruct struct {
		testName      string
		flag          string
		file          string
		expectedCount int64
	}
	fillStruct := func (name, flag, file string) TestStruct {
		filePath := getFilePath(file)
		return TestStruct{
			testName: name,
			flag: flag,
			file: filePath,
			expectedCount: countWithWC(flag, filePath),
		}
	}

	countTests := []TestStruct{
		fillStruct("byte count", "-c", "test.txt"),
		fillStruct("line count", "-l", "test.txt"),
	}

	for _, test := range countTests {
		t.Run(test.testName, func(t *testing.T) {
			got := WordCounter(test.flag, test.file)
			if got != test.expectedCount {
				t.Errorf("got: %v, want: %v\n", got, test.expectedCount)
			}
			fmt.Printf("got: %v, want: %v\n", got, test.expectedCount)
		})
	}
}