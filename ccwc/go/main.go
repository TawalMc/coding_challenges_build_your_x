package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// defer duration(track("main"))

	lineCount := flag.Bool("l", false, "count number of line in a file")
	wordCount := flag.Bool("w", false, "count number of word in a file")
	charCount := flag.Bool("m", false, "count number of character in a file")
	byteCount := flag.Bool("c", false, "count number of byte in a file")

	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		log.Fatal("provide a file(s) path(s)")
	}
	// fmt.Println(files)

	args := CWArgs{
		l: ArgAndCount{*lineCount, 0},
		w: ArgAndCount{*wordCount, 0},
		m: ArgAndCount{*charCount, 0},
		c: ArgAndCount{*byteCount, 0},
	}

	if !args.l.arg && !args.w.arg &&
		!args.m.arg && !args.c.arg {
		args = CWArgs{
			l: ArgAndCount{true, 0},
			w: ArgAndCount{true, 0},
			m: ArgAndCount{true, 0},
			c: ArgAndCount{true, 0},
		}
	}

	resultChannel := make(chan ResultChan)

	for _, file := range files {
		go func() {
			argsAndCounts, err := WordCounter(args, file)
			if err != nil {
				log.Fatal(err)
			}
			resultChannel <- ResultChan{file, argsAndCounts}
		}()
	}

	multipleArgsFiles := CWArgs{
		l: ArgAndCount{*lineCount, 0},
		w: ArgAndCount{*wordCount, 0},
		m: ArgAndCount{*charCount, 0},
		c: ArgAndCount{*byteCount, 0},
	}
	for idx := 0; idx < len(files); idx++ {
		r := <-resultChannel
		fmt.Println(r.cw, r.f)

		// if len(files) != 0 {
		multipleArgsFiles.l.count += r.cw.l.count
		multipleArgsFiles.w.count += r.cw.w.count
		multipleArgsFiles.m.count += r.cw.m.count
		multipleArgsFiles.c.count += r.cw.c.count
		// }
	}
	
	if len(files) > 1 {
		fmt.Println(multipleArgsFiles, files)
	}

}

func PrintDefaults() {
	panic("unimplemented")
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}
