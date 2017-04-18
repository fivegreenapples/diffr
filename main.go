package main

import (
	"bufio"
	"crypto/md5" // #nosec
	"fmt"
	"io"
	"os"
)

type md5Slice [][md5.Size]byte
type lineSlice [][]byte

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "error: missing input files")
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		fmt.Fprintln(os.Stderr, "error: missing input file")
		os.Exit(1)
	}
	if len(os.Args) >= 4 {
		fmt.Fprintln(os.Stderr, "error: too many arguments")
		os.Exit(1)
	}

	fileA, errA := os.Open(os.Args[1])
	if errA != nil {
		fmt.Fprintf(os.Stderr, "error opening '%s': %s\n", os.Args[1], errA)
	}
	fileB, errB := os.Open(os.Args[2])
	if errB != nil {
		fmt.Fprintf(os.Stderr, "error opening '%s': %s\n", os.Args[2], errB)
	}
	if errA != nil || errB != nil {
		os.Exit(2)
	}

	hashesA, linesA, err := hashLines(fileA)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
	hashesB, linesB, err := hashLines(fileB)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}

	lcs := findLCS(hashesA, hashesB)

	cursorA, cursorB, cursorCom := 0, 0, 0
	for cursorCom < len(lcs) || cursorA < len(hashesA) || cursorB < len(hashesB) {
		for cursorCom < len(lcs) &&
			cursorA < len(hashesA) &&
			cursorB < len(hashesB) &&
			lcs[cursorCom] == hashesA[cursorA] &&
			lcs[cursorCom] == hashesB[cursorB] {
			// common item, move on all cursors
			cursorCom++
			cursorA++
			cursorB++
		}
		for cursorA < len(hashesA) &&
			(cursorCom >= len(lcs) || lcs[cursorCom] != hashesA[cursorA]) {
			fmt.Printf("- %s", linesA[cursorA])
			cursorA++
		}
		for cursorB < len(hashesB) &&
			(cursorCom >= len(lcs) || lcs[cursorCom] != hashesB[cursorB]) {
			fmt.Printf("+ %s", linesB[cursorB])
			cursorB++
		}
	}

}

func hashLines(r io.Reader) (md5Slice, lineSlice, error) {
	lines := lineSlice{}
	hashes := md5Slice{}
	bufR := bufio.NewReader(r)

	var err error
	var line []byte
	for err == nil {
		line, err = bufR.ReadBytes('\n')
		if len(line) > 0 {
			lines = append(lines, line)
			hashes = append(hashes, md5.Sum(line)) // #nosec
		}

	}

	if err != io.EOF {
		return md5Slice{}, lineSlice{}, err
	}

	return hashes, lines, nil
}
