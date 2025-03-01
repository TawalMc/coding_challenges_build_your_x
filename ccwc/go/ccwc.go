package main

import (
	"bufio"
	"io"
	"os"
)

const (
	MAX_BYTES_READ = 1024 * 5
)

func WordCounter(cwArgs CWArgs, file string) (CWArgs, error) {
	args := cwArgs
	// defer duration(track("wc"))

	f, err := os.Open(file)
	if err != nil {
		return CWArgs{}, err
	}

	if args.l.arg {
		args.l.count, err = lineCount(f)

		if err != nil {
			return CWArgs{}, err
		}
	}

	if args.w.arg {
		args.w.count, err = wordCount(f)

		if err != nil {
			return CWArgs{}, err
		}
	}

	if args.m.arg {
		args.m.count, err = charCount(f)

		if err != nil {
			return CWArgs{}, err
		}
	}

	if args.c.arg {
		args.c.count, err = byteCountV1(f)

		if err != nil {
			return CWArgs{}, err
		}
	}

	f.Close()
	return args, nil
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
