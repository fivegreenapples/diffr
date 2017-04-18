package main

/**
 * Find the Longest Common Subsequence
 * @see https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
 */

func findLCS(A, B md5Slice) md5Slice {
	if len(A) == 0 || len(B) == 0 {
		return md5Slice{}
	}

	startIndex := findStartIndex(A, B)
	endIndexA, endIndexB := findEndIndices(A, B, startIndex)

	if startIndex >= endIndexA {
		return A
	}
	if startIndex >= endIndexB {
		return B
	}

	partialLCS := doLCS(A[startIndex:endIndexA], B[startIndex:endIndexB])

	lcs := make(md5Slice, startIndex)
	copy(lcs, A[:startIndex])
	lcs = append(lcs, partialLCS...)
	lcs = append(lcs, A[endIndexA:]...)
	return lcs

}

func findStartIndex(A, B md5Slice) int {
	start, i, j := 0, 0, 0
	for i < len(A) && j < len(B) && A[i] == B[j] {
		start = i + 1
		i++
		j++
	}
	return start
}

func findEndIndices(A, B md5Slice, startIndex int) (int, int) {
	endA, endB := len(A), len(B)
	for endA > startIndex && endB > startIndex && endA > 0 && endB > 0 && A[endA-1] == B[endB-1] {
		endA--
		endB--
	}
	return endA, endB
}

func doLCS(A, B md5Slice) md5Slice {

	partialsLength := len(B) + 1
	partials := make([]md5Slice, partialsLength)

	for _, elA := range A {
		newPartials := make([]md5Slice, partialsLength)
		for iB, elB := range B {
			if elA == elB {
				newPartials[iB+1] = partials[iB]
				newPartials[iB+1] = append(newPartials[iB+1], elA)
			} else {
				if len(partials[iB+1]) > len(newPartials[iB]) {
					newPartials[iB+1] = partials[iB+1]
				} else {
					newPartials[iB+1] = newPartials[iB]
				}
			}
		}
		partials = newPartials
	}
	return partials[partialsLength-1]
}
