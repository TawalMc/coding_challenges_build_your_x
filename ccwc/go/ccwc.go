package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	MAX_BYTES_READ = 1024 * 5
)

type ArgAndCount struct {
	arg   bool
	count int64
}

type CWArgs struct {
	l ArgAndCount
	w ArgAndCount
	m ArgAndCount
	c ArgAndCount
}

func WordCounter(args *CWArgs, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	if args.l.arg {
		args.l.count, err = lineCount(f)

		if err != nil {
			return err
		}
	}

	if args.w.arg {
		args.w.count, err = wordCount(f)

		if err != nil {
			return err
		}
	}

	if args.m.arg {
		args.m.count, err = charCount(f)

		if err != nil {
			return err
		}
	}

	if args.c.arg {
		args.c.count, err = byteCountV1(f)

		if err != nil {
			return err
		}
	}

	f.Close()
	return nil
}

func lineCount(file *os.File) (int64, error) {
	// defer duration(track("lineCount"))
	file.Seek(0, io.SeekStart)

	var count int64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func wordCount(file *os.File) (int64, error) {
	// defer duration(track("wordCount"))
	file.Seek(0, io.SeekStart)

	var count int64 = 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func charCount(file *os.File) (int64, error) {
	// defer duration(track("charCount"))
	file.Seek(0, io.SeekStart)

	var count int64 = 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func byteCountV1(file *os.File) (int64, error) {
	// defer duration(track("byteCountV1"))
	file.Seek(0, io.SeekStart)

	var count int64 = 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

// check after: reflection
func formatOuput(args CWArgs) string {
	counts := make([]string, 4)
	idx := 0

	if args.l.arg {
		counts[idx] = "l: " + fmt.Sprint(args.l.count)
		idx++
	}

	if args.w.arg {
		counts[idx] = "w: " + fmt.Sprint(args.w.count)
		idx++
	}

	if args.m.arg {
		counts[idx] = "m: " + fmt.Sprint(args.m.count)
		idx++
	}

	if args.c.arg {
		counts[idx] = "c: " + fmt.Sprint(args.c.count)
		idx++
	}

	if idx == 0 {
		return ""
	}

	output := ""
	for i := 0; i < idx-1; i++ {
		output += counts[i] + ", "
	}
	output += counts[idx-1]

	return output
}

/* func byteCountV2(file *os.File) (int64, error) {
	// defer duration(track("byteCountV2"))

	var count int64 = 0
	byteContainer := make([]byte, MAX_BYTES_READ)
	for {
		read, err := file.Read(byteContainer)
		if err != nil && err != io.EOF {
			return 0, err
		}

		count += int64(read)
		if read < MAX_BYTES_READ {
			break
		}
	}
	return count, nil
}
*/
