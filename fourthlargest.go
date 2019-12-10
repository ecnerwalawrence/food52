package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

/*******

PROBLEM:

Write a function that takes a list of integers and returns the 4 highest in O(n) time.
We'd like to see how you think about algorithms and data structures. So please use your own logic instead of calling out to libraries.

******/

/***

Thought process of towards the solution.
- Above did not state the list is sorted.  Thus, I cannot use an O(logN) solution.
- The problem is limited to a constant (4th) largest.  Therefore, sorting is unnecessary, which cost O(NlogN) for the best solution.
- Alternate Solution:
	* We iterate the loop once.
	* For each number we compare 4 numbers for the largest 4 numbers.
	The order of operation is still the same.  Loop is N and each loop we compare about 4 times to maintain the largest 4th.  Thus, 4N operations.
	Also, I think code complexity is about the same.

***/

// Largest - finds the largest and returns the index
//           I split this logic to breakdown the problem and simplify main algorithm.
//	args:
//		nums - list of numbers
//	Returns:
//		index of the largest
func Largest(nums []int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	largestIdx := 0
	for i := 1; i < l; i++ {
		if nums[i] > nums[largestIdx] {
			largestIdx = i
		}
	}
	return largestIdx
}

// FourthLargest - finds the 4th largest and returns the 4th largest number
//	args:
//		nums - list of numbers
//	Returns:
//		4th largest
func FourthLargest(nums []int) (int, error) {
	// Preserving nums
	if len(nums) == 0 {
		return 0, errors.New("empty list")
	}
	numsToSearch := make([]int, len(nums))
	copy(numsToSearch, nums)

	// 4N operation = O(N)
	largestIdx := -1
	nthLargest := nums[0]
	for n := 0; n < 4; n++ {
		largestIdx = Largest(numsToSearch)
		if largestIdx < 0 {
			break
		}
		nthLargest = numsToSearch[largestIdx]
		// The list is unsorted so we do not need to maintain order.
		// replace the largest with e
		endIdx := len(numsToSearch) - 1
		numsToSearch[largestIdx] = numsToSearch[endIdx]
		numsToSearch = numsToSearch[:endIdx]
	}
	return nthLargest, nil
}

func main() {

	jsonPtr := flag.String("json", "sample_input.json", "a string")
	flag.Parse()

	byt, err := ioutil.ReadFile(*jsonPtr)
	if err != nil {
		log.Fatal(err)
	}

	testSets := [][]int{}
	if err := json.Unmarshal(byt, &testSets); err != nil {
		log.Fatal(err)
	}

	for i, nums := range testSets {
		log.Println("TEST #", i)
		log.Println(nums)
		fourthLargest, err := FourthLargest(nums)
		if err != nil {
			log.Println("error:", err.Error())
			continue
		}
		log.Println("result:", fourthLargest)
	}
}
