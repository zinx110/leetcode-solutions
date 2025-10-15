package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type TrieNode struct {
	children  []*TrieNode
	endOfWord bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}

func (w *WordDictionary) addWord(word string) {

}

func (w *WordDictionary) search(word string) bool {
	return true

}

func main() {
	file := "211. Design Add and Search Words Data Structure/testcase.json"
	path := filepath.Join(file)
	jsonData, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading json data : %v", err)
		return
	}
	var testcase TestCase
	err = json.Unmarshal(jsonData, &testcase)
	if err != nil {
		fmt.Printf("Error unmarshalling json data : %v", err)
		return
	}

	passNum := 0
	var wordDictionary WordDictionary

	for _, test := range testcase.Tests {

		output := []any{}
		for i := 0; i < len(test.Input.Operations); i++ {
			res := any(nil)
			switch test.Input.Operations[i] {
			case "WordDictionary":
				wordDictionary = Constructor()
			case "addWord":
				wordDictionary.addWord(test.Input.Arguments[i][0])
			case "search":
				res = wordDictionary.search(test.Input.Arguments[i][0])

			}
			output = append(output, res)

		}
		if reflect.DeepEqual(output, test.ExpectedOutput) {
			passNum++
			fmt.Println("Test passed")
		} else {
			fmt.Println("Test failed")
		}
		fmt.Printf("operations        : %v\n", test.Input.Operations)
		fmt.Printf("arguments         : %v\n", test.Input.Arguments)
		fmt.Printf("output            : %v\n", output)
		fmt.Printf("expected ouput    : %v\n", test.ExpectedOutput)
		fmt.Println("---------------")
	}
	fmt.Printf("Total tests : %v, Passed: %v, Failed: %v", len(testcase.Tests), passNum, len(testcase.Tests)-passNum)
}

type TestCase struct {
	Tests []Test `json:"tests"`
}
type Test struct {
	Input          Input `json:"input"`
	ExpectedOutput []any `json:"expectedOutput"`
}
type Input struct {
	Operations []string   `json:"operations"`
	Arguments  [][]string `json:"arguments"`
}
