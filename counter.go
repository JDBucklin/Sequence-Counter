package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Count holds the number of occurences of a given set of words
type Count struct {
	words string
	count int
}

func main() {
	// check if stdin has input or if args were passed to the program
	// more info about stat.Mode() here https://golang.org/pkg/os/#FileInfo
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal("error checking input on stdin")
	}
	stdinHasInput := stat.Mode()&os.ModeCharDevice == 0
	hasArgs := len(os.Args[1:]) != 0
	if !hasArgs && !stdinHasInput {
		fmt.Println("exiting must pass file names or send data on stdin")
		fmt.Println("example usage: ./newrelic file1.txt file2.txt ...")
		fmt.Println("or")
		fmt.Println("example usage: cat file1.txt | ./newrelic")
		return
	}

	// read contents from stdin
	if stdinHasInput {
		counts := getCounts(os.Stdin, "stdin")
		orderedCounts := getOrderedCounts(counts)
		printCounts("stdin", orderedCounts)
	}

	// read contents of each file passed as an arg
	if hasArgs {
		args := os.Args[1:]
		for _, arg := range args {
			file, err := os.Open(arg)
			if err != nil {
				log.Fatalf("error reading file %s", arg)
			}
			defer file.Close()

			counts := getCounts(file, arg)
			orderedCounts := getOrderedCounts(counts)
			printCounts(arg, orderedCounts)
		}
	}
}

// getCounts searches through the given file line by line and creates a map
// of all three word sequences and the number of times each occurs
func getCounts(file *os.File, filename string) map[string]int {
	rx := regexp.MustCompile(`\w+'?(\w+)?`)
	counts := map[string]int{}
	scanner := bufio.NewScanner(file)
	var previous1, previous2 string

	for scanner.Scan() {
		words := rx.FindAllString(scanner.Text(), -1)
		for _, word := range words {
			word = strings.ToLower(word)
			if previous1 != "" && previous2 != "" {
				key := fmt.Sprintf("%s %s %s", previous2, previous1, word)
				val := counts[key]
				val++ // if the key doesn't exist yet val will be 0
				counts[key] = val
			}
			previous2 = previous1
			previous1 = word
		}
	}
	if scanner.Err() != nil {
		log.Fatalf("error scanning file %s %s", filename, scanner.Err())
	}
	return counts
}

// getOrderedCounts copies the contents of the passed in counts map into a slice
// of Count objects. It then sorts the slice by Count.count and then alphabetically
// by the three word sequence stored in Count.words
func getOrderedCounts(counts map[string]int) []Count {
	orderedCounts := []Count{}
	for key, value := range counts {
		orderedCounts = append(orderedCounts, Count{key, value})
	}

	sort.Slice(orderedCounts, func(i, j int) bool {
		if orderedCounts[i].count > orderedCounts[j].count {
			return true
		}
		if orderedCounts[i].count < orderedCounts[j].count {
			return false
		}
		return orderedCounts[i].words < orderedCounts[j].words
	})
	return orderedCounts
}

// printCounts prints out the top 100 three word sequences
// note that anything after 100 entries will not be printed even if the number of
// occurrences matches the lowest count printed
func printCounts(filename string, orderedCounts []Count) {
	fmt.Printf("Three word sequence counts for %s:\n", filename)
	for i := 0; i < 100 && i < len(orderedCounts); i++ {
		fmt.Printf("%d: %s - %d\n", i+1, orderedCounts[i].words, orderedCounts[i].count)
	}
	fmt.Println()
}
