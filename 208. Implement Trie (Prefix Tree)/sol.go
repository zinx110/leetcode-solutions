package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}

}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, c := range word {
		if _, ok := curr.children[c]; !ok {
			curr.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		curr = curr.children[c]
	}
	curr.endOfWord = true

}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, c := range word {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return curr.endOfWord

}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, c := range prefix {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return true

}

func main() {
	// currPath, _ := os.Getwd()
	// fmt.Println("Current path: ", currPath)

	filename := "./208. Implement Trie (Prefix Tree)/testcase.json"
	path := filepath.Join(filename)
	// fmt.Printf("Filepath: %v", path)

	jsonData, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading file: %v", err)
		return
	}

	// fmt.Printf("File data: %v, type %T", string(jsonData), jsonData)
	var testCase TestCase
	err = json.Unmarshal(jsonData, &testCase)
	if err != nil {
		fmt.Printf("error unmarshalling file: %v", err)
		return
	}
	passNum := 0

	for _, test := range testCase.Tests {
		trie := Constructor()
		input := test.Input
		commands := input.Commands
		arguments := input.Arguments
		expectedOutput := test.ExpectedOutput
		output := []any{}
		for i := 0; i < len(commands); i++ {
			res := any(nil)
			if commands[i] == "Trie" {
				trie = Constructor()
			} else if commands[i] == "insert" {
				arg := arguments[i][0]
				trie.Insert(arg)
			} else if commands[i] == "search" {
				arg := arguments[i][0]
				res = trie.Search(arg)
			} else if commands[i] == "startsWith" {
				arg := arguments[i][0]
				res = trie.StartsWith(arg)
			}
			output = append(output, res)
		}
		if !reflect.DeepEqual(output, expectedOutput) {
			fmt.Printf("Test failed")
		} else {
			passNum++
			fmt.Printf("Test passed")
		}
		fmt.Printf("\nInput: \n  Arguments: %v \n  Commands: %v \n  Output: %v \n  Expected Output: %v \n", arguments, commands, output, expectedOutput)
		fmt.Println("--------------------------------")
	}
	fmt.Printf("Total tests: %v, Passed tests: %v, Failed tests: %v", len(testCase.Tests), passNum, len(testCase.Tests)-passNum)
}

type Input struct {
	Commands  []string   `json:"commands"`
	Arguments [][]string `json:"arguments"`
}
type Test struct {
	Input          Input `json:"input"`
	ExpectedOutput []any `json:"expectedOutput"`
}
type TestCase struct {
	Tests []Test `json:"tests"`
}
