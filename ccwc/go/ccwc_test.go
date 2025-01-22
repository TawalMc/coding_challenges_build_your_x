package main

import (
	"fmt"
	"os"
	"path/filepath"

	// "os/exec"
	"testing"
)

func TestWordCounter(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf("error current dir: %v", err)
	}
	subFolder := "../"
	fileName := "test.txt"

	testFilePath, err := filepath.Abs(filepath.Join(cwd, subFolder, fileName))

	if err != nil {
		t.Errorf("error abs path: %v", err)
	}
	
	t.Run("byte count", func(t *testing.T) {
		// fileTestDir, err := os.Getwd()

		fmt.Println("yeah", testFilePath)

		// want := 0
		// got := WordCounter("")

		// if got != want {

		// }

		// cmd := exec.Command("wc", "-c", "")
	})
}
