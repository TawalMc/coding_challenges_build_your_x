package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func WordCounterFile(cwArgs CWArgs, file string) (CWArgs, error) {
	f, err := os.Open(file)
	if err != nil {
		return CWArgs{}, err
	}

	return WordCounter(cwArgs, f)
}

func WordCounter(cwArgs CWArgs, file *os.File) (CWArgs, error) {
	// defer duration(track("wc"))
	args := cwArgs

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return CWArgs{}, err
		}

		if args.l.arg {
			args.l.count++
		}

		if args.w.arg {
			wCount, err := wordCount(line)
			args.w.count += wCount

			if err != nil {
				return CWArgs{}, err
			}
		}

		if args.m.arg {
			mCount, err := charCount(line)
			args.m.count += mCount

			if err != nil {
				return CWArgs{}, err
			}
		}

		if args.c.arg {
			cCount, err := byteCountV1(line)
			args.c.count += cCount

			if err != nil {
				return CWArgs{}, err
			}
		}

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return args, nil
}

func wordCount(text string) (int64, error) {
	// defer duration(track("wordCount"))
	content := strings.NewReader(text)
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanWords)

	var count int64 = 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func charCount(text string) (int64, error) {
	// defer duration(track("charCount"))
	content := strings.NewReader(text)
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanRunes)

	var count int64 = 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func byteCountV1(text string) (int64, error) {
	// defer duration(track("byteCountV1"))
	content := strings.NewReader(text)
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanBytes)

	var count int64 = 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
