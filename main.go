package main

import (
	"bufio" // #nosec
	"fmt"
	"io"
	"os"

	"github.com/fivegreenapples/diff"
)

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

	linesA, err := readLines(fileA)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
	linesB, err := readLines(fileB)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}

	patch := diff.MakeStringPatch(linesA, linesB)

	indexA := 0
	indexB := 0
	indexP := 0

	for indexP < len(patch) || indexA < len(linesA) {
		if indexP < len(patch) {
			currentChange := patch[indexP]
			if currentChange.Offset == indexA {
				// Act on current change as we're at the right line number

				// First render the header
				if len(currentChange.Add) > 0 && currentChange.Skip > 0 {
					// A modification
					fmt.Printf("%d", indexA+1)
					if currentChange.Skip > 1 {
						fmt.Printf(",%d", indexA+currentChange.Skip)
					}
					fmt.Print("c")
					fmt.Printf("%d", indexB+1)
					if len(currentChange.Add) > 1 {
						fmt.Printf(",%d", indexB+len(currentChange.Add))
					}
					fmt.Print("\n")
				} else if currentChange.Skip > 0 {
					// A deletion
					fmt.Printf("%d", indexA+1)
					if currentChange.Skip > 1 {
						fmt.Printf(",%d", indexA+currentChange.Skip)
					}
					fmt.Printf("d%d\n", indexB)
				} else if len(currentChange.Add) > 0 {
					// An addition
					fmt.Printf("%da%d", indexA, indexB+1)
					if len(currentChange.Add) > 1 {
						fmt.Printf(",%d", indexB+len(currentChange.Add))
					}
					fmt.Print("\n")
				}

				if currentChange.Skip > 0 {
					// A deletion of lines
					for s := 0; s < currentChange.Skip; s++ {
						fmt.Printf("< %s", linesA[indexA])
						indexA++
					}
				}
				if currentChange.Skip > 0 && len(currentChange.Add) > 0 {
					fmt.Print("---\n")
				}
				if len(currentChange.Add) > 0 {
					// An addition of lines
					for _, l := range currentChange.Add {
						fmt.Printf("> %s", l)
						indexB++
					}
				}
				indexP++
				continue
			} else if indexA >= len(linesA) {
				// this protects us from a duff patch where items in the patch
				// reference beyond the end of the original.
				break
			}
		}

		if indexA < len(linesA) {
			indexA++
			indexB++
		}
	}

}

func readLines(r io.Reader) ([]string, error) {
	lines := []string{}
	bufR := bufio.NewReader(r)

	var err error
	var line []byte
	for err == nil {
		line, err = bufR.ReadBytes('\n')
		if len(line) > 0 {
			lines = append(lines, string(line))
		}

	}

	if err != io.EOF {
		return []string{}, err
	}

	return lines, nil
}
