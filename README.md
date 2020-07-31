# Sequence-Counter
This program takes text as input and outputs the top 100 reoccuring three word sequences and the number of times they occur.\
All punctuation and newlines are ignored, however, both letters and numbers are included.

# Usage
Build the project using: `go build counter.go`\
You can run the executable using one of the following methods:\
Passing a list of files as args `./counter file1.txt file2.txt ...`\
Passing input using stdin in `echo "one two three four" | ./counter`\
Using a combination of the two previous methods `echo "one two three four" | ./counter file1.txt file2.txt ...`

The output will list the file name (or stdin) and the counts for the top 100 three word sequences.
```
Three word sequence counts for stdin:
1: four five six - 1
2: one two three - 1
3: three four five - 1
4: two three four - 1

Three word sequence counts for file1.txt:
...
```