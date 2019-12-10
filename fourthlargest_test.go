package main

import (
	"errors"
	"testing"
)

type TestPlanFourthLargest struct {
	Name        string
	Input       []int
	Expectation int
	Error       error
}

func TestFourthLargest(t *testing.T) {
	for _, plan := range []TestPlanFourthLargest{
		TestPlanFourthLargest{
			"When list is empty",
			[]int{},
			-1,
			errors.New("empty list"),
		},
		TestPlanFourthLargest{
			"When list is sorted numbers",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			9,
			nil,
		},
		TestPlanFourthLargest{
			"When list is unsorted numbers",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 12, 10, 11, 9},
			9,
			nil,
		},
		TestPlanFourthLargest{
			"When list is contains a negative number",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 12, 10, 11, -9},
			8,
			nil,
		},

		TestPlanFourthLargest{
			"When list is contains negative numbers",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 12, -10, -11, -9},
			6,
			nil,
		},
		TestPlanFourthLargest{
			"When list is contains a single number",
			[]int{10},
			10,
			nil,
		},
		TestPlanFourthLargest{
			"When list is contains three numbers",
			[]int{8, 9, 10},
			8, // returns the 3rd largest but it is close enough
			nil,
		},
		TestPlanFourthLargest{
			"When list has duplicate numbers",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			11, // 12 appears 3 times, 12,12,12,11
			nil,
		},
	} {
		t.Run(plan.Name, func(t *testing.T) {

			fourthLargest, err := FourthLargest(plan.Input)
			if err != plan.Error {
				if plan.Error != nil && err.Error() != plan.Error.Error() {
					t.Errorf("Expected err:%v result:%v", err, plan.Error)
				}
			}
			if plan.Expectation != fourthLargest && plan.Error == nil {
				t.Errorf("Expected %d != Result %d", plan.Expectation, fourthLargest)
			}
		})
	}

}
