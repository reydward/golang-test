package level3

import "testing"

type testCase struct {
	from int
	to int
	result []int
}

func TestFilterByAgeRange(t *testing.T) {
	ages := []int{1,2,3,4,5,6,7,8,9}

	testCases := []testCase{
		{1, 2, []int{1,2}},
		{5, 9, []int{5,6,7,8,9}},
		{3, 1, []int{}},
	}

	for _, t2 := range testCases {
		result := filterByAgeRange(t2.from, t2.to, ages)
		if !Equal(t2.result, result) {
			t.Error("This is not the expected result.")
		}
	}

}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}