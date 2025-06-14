ccwc is a rewrite version of wc: word, line, character, and byte count

Description: [Build Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc/)

### Toolchain

- Go version: go1.23.4 linux/amd64
- OS: Ubuntu 22.04. LTS

### To test

#### Build from source
```bash
git clone https://github.com/TawalMc/coding_challenges_build_your_x.git
cd coding_challenges_build_your_x/ccwc/go
```
```
go build -o ccwc .
./ccwc /path/to/file
```

#### Get release from Github (ccwc_go (ccwc written in Go))
Download **ccwc** [here](https://github.com/TawalMc/coding_challenges_build_your_x/releases/tag/ccwc_go_v1.0.0)

They are some points about this Golang implementation of wc (or ccwc):

### Format to print counts
I print the counts in the following format:
```
l: <lines>, w: <words>, m: <characters>, c: <bytes> <file>
```
#### example
```
l: 111, w:100, m: 200, c: 200 text.txt  
```
In case of multiples files, the print is the same but each file with its line
followed by a line with the sums of each type of counts 

#### example
```
l: 111, w:100, m: 200, c: 200 text.txt  
l: 111, w:100, m: 200, c: 200 content.txt
l: 222, w:200, m: 400, c: 400 [text.txt content.txt]  
```

### Handle multiple files
I use Go concurreny to handle multiple files arguments. So in this case, I did iteration on the files, then call the main function __WordCounter__  in goroutine. This make the program fast if there are multiple files
