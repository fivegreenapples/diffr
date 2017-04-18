# diffr

This is a pretty basic diff tool that I created to play with the Longest Common Subsequence algorithm. See [wikipedia](https://en.wikipedia.org/wiki/Longest_common_subsequence_problem) for details.

Usage:

```
diffr file1.txt file2.txt
```

Example output:

```
+ this line was added
+ so was this line
- but this line was taken away
```

You'll note the output doesn't give anything helpful like line numbers or instructions for `patch`. It is that basic...
