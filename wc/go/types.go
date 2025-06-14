package main

import "fmt"

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

func (args CWArgs) String() string {
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

type ResultChan struct {
	f  string
	cw CWArgs
}
