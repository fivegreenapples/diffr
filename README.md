# diffr

This is a pretty basic diff tool that I created to play with the Longest Common Subsequence algorithm. See [wikipedia](https://en.wikipedia.org/wiki/Longest_common_subsequence_problem) for details. Subsequent to the initial commit of this tool, I moved the meat of the algorithm to a separate repo: (diff)[https://github.com/fivegreenapples/diff] for use in other projects.

Usage:

```
diffr file1.txt file2.txt
```

Output mimics the output from the standard diff tool. e.g.:

```
1a2
> This line was added
3,4d3
< This line was deleted
< So was this line
6,7c5,6
< These lines were changed...
< ...into something else
---
> They were changed into...
> ...much nicer lines
```

