package main

import (
	"os"
	"reflect"
	"testing"
)

func TestGetCounts(t *testing.T) {
	type GetCountsData struct {
		filename string
		expected map[string]int
	}

	dataItems := []GetCountsData{
		{"./testdata/sample1.txt",
			map[string]int{
				"six six six":      9,
				"one two three":    2,
				"two three four":   2,
				"three four five":  2,
				"four five six":    2,
				"five six seven's": 2,
				"six seven's one":  1,
				"seven's one two":  1,
				"six seven's six":  1,
				"seven's six six":  1,
			},
		},
		{"./testdata/sample2.txt",
			map[string]int{
				"i love sandwiches": 2,
				"love sandwiches i": 1,
				"sandwiches i love": 1,
			},
		},
		{"./testdata/sample3.txt",
			map[string]int{
				"echo1 1 echo2": 1,
				"1 echo2 2":     1,
				"echo2 2 echo3": 1,
				"2 echo3 3":     1,
			},
		},
	}

	for _, item := range dataItems {
		file, err := os.Open(item.filename)
		if err != nil {
			t.Fatal(err)
		}
		result := getCounts(file, item.filename)
		if reflect.DeepEqual(result, item.expected) == false {
			t.Fatalf("failed on sample1.txt expected %v got %v", item.expected, result)
		}
	}
}

func TestGetOrderedCounts(t *testing.T) {
	type GetOrderedCountsData struct {
		input    map[string]int
		expected []Count
	}

	dataItems := []GetOrderedCountsData{
		{map[string]int{
			"one two three":  2,
			"two three four": 1,
			"too three four": 1,
			"seven six six":  9,
		},
			[]Count{
				Count{"seven six six", 9},
				Count{"one two three", 2},
				Count{"too three four", 1},
				Count{"two three four", 1},
			},
		},
		{map[string]int{}, []Count{}},
	}

	for _, item := range dataItems {
		result := getOrderedCounts(item.input)
		if reflect.DeepEqual(result, item.expected) == false {
			t.Fatalf("failed expected %v got %v", item.expected, result)
		}
	}
}
