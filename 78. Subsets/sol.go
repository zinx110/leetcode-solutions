package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func subsets(nums []int) [][]int {
	res := [][]int{make([]int, 0)}
	curr := make([]int, 0, len(nums))
	var dfs func(j int)
	dfs = func(j int) {
		if j > len(nums) {
			return
		}
		for i := j; i < len(nums); i++ {
			curr = append(curr, nums[i])
			dfs(i + 1)
			copiedCurr := make([]int, len(curr))
			copy(copiedCurr, curr)
			res = append(res, copiedCurr)
			curr = curr[:len(curr)-1]
		}
	}
	dfs(0)
	return res
}
func subsetsBFS(nums []int) [][]int {
	res := [][]int{make([]int, 0)}
	for _, num := range nums {
		currentRes := make([][]int, 0, len(res)*2)
		for _, subset := range res {
			currentRes = append(currentRes, append(subset, num))
			currentRes = append(currentRes, subset)
		}
		res = currentRes
	}
	return res
}

func main() {
	currentPath := "./78. Subsets"
	testCaseFile := "testcase.json"
	file := filepath.Join(currentPath, testCaseFile)
	testcaseData, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file data")
		return
	}
	var testcases Testcases
	err = json.Unmarshal(testcaseData, &testcases)
	if err != nil {
		fmt.Println("Error unmarshalling json data")
		return
	}

	passNum := 0
	for _, test := range testcases.Tests {

		input := test.Input.Nums
		expectedOutput := test.ExpectedOutput
		output := subsetsBFS(input)
		outputKeys := cannonicalizeSubsets(output)
		expectedKeys := cannonicalizeSubsets(expectedOutput)
		if reflect.DeepEqual(outputKeys, expectedKeys) {
			fmt.Println("Test Passed")
			passNum++
		} else {
			fmt.Println("Test Failed")
		}
		fmt.Printf(
			"Input           : %v\nOutput          : %v\nExpected Output : %v\n-----------------\n\n",
			input, output, expectedOutput,
		)
	}

	fmt.Printf("Total tests: %v, Passed: %v, Failed %v", len(testcases.Tests), passNum, len(testcases.Tests)-passNum)

}

type Testcases struct {
	Tests []Test `json:"tests"`
}

type Test struct {
	Input          Input   `json:"input"`
	ExpectedOutput [][]int `json:"expectedOutput"`
}
type Input struct {
	Nums []int `json:"nums"`
}

func cannonicalizeSubsets(sets [][]int) []string {
	keys := make([]string, len(sets))
	for i, subset := range sets {
		copied := append([]int(nil), subset...)
		sort.Ints(copied)
		var b strings.Builder
		for j, v := range copied {
			if j > 0 {
				b.WriteByte(',')
			}
			b.WriteString(strconv.Itoa(v))
		}
		keys[i] = b.String()
	}
	sort.Strings(keys)
	return keys
}
