package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
)

func main() {
	fileData, err := os.Open("./data.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bufio.NewReader(fileData))
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF {
			if i := bytes.IndexByte(data, ','); i >= 0 {
				return i + 1, data[0:i], nil
			}
		}
		return 0, nil, nil
	})
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	currentDir, _ := os.Getwd()
	absPath := path.Join(currentDir, "latin.txt")
	dat, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))
	fmt.Println("reading buffered step by step")
	outFile := path.Join(currentDir, "resultfile.txt")
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	infile, _ := os.Open(absPath)
	writer := bufio.NewWriter(f)
	reader := bufio.NewReader(infile)
	currentBuffer := make([]byte, 1024)
	for {
		currentBytes, err := reader.Read(currentBuffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err != nil && err == io.EOF {
			break
		}
		fmt.Printf("Read %d \n", currentBytes)
		writer.Write(currentBuffer)
	}
	writer.Flush()
	f.Close()
}
