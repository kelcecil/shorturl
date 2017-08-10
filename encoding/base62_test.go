package encoding

import "testing"

func TestHashToID(t *testing.T) {
	testCases := map[string]int{
		"bM":  100,
		"ba":  62,
		"bb":  63,
		"cLs": 10000,
	}

	for id, desiredTestResult := range testCases {
		result := HashToID(id)
		if result != desiredTestResult {
			t.Errorf("hash was not converted back to ID successfully. Expected: %v, Got %v", desiredTestResult, result)
		}
	}

}

// TestIdToHash ... Test that identifier numbers are properly converted into
// short ids properly.
func TestIDToHash(t *testing.T) {
	testCases := map[int]string{
		100:   "bM",
		62:    "ba",
		63:    "bb",
		10000: "cLs",
	}

	for testID, desiredTestResult := range testCases {
		result := IDToHash(testID)
		if result != desiredTestResult {
			t.Errorf("id to shortened string failed for value: %v; got: %v", testID, result)
		}
	}
}

func TestMapDigitsToAlphabet(t *testing.T) {
	testCases := map[string][]int{
		"a":   []int{0},
		"b":   []int{1},
		"aa":  []int{0, 0},
		"ab":  []int{0, 1},
		"abc": []int{0, 1, 2},
		"ba":  []int{1, 0},
		"Mb":  []int{38, 1},
	}
	for expectedString, digits := range testCases {
		result := MapDigitsToAlphabet(digits)
		if result != expectedString {
			t.Errorf("Result did not match expected. Got: %v. Expected: %v", result, expectedString)
		}
	}
}

// TestDigitsForInt ... Ensure that identifiers are broken down into
// base62 parts correctly.
func TestDigitsForInt(t *testing.T) {
	testCases := map[int][]int{
		100:   []int{1, 38},
		62:    []int{1, 0},
		10000: []int{2, 37, 18},
	}

	for testID, desiredTestResult := range testCases {
		computedResult := FindDigitsForInt(testID)
		if !slicesAreEqual(computedResult, desiredTestResult) {
			t.Errorf("Converting id to base 62 failed for value: %v; got: %v", testID, computedResult)
		}
	}
}

// slicesAreEqual ... Helper for comparing two slices of integers for equality.
func slicesAreEqual(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
