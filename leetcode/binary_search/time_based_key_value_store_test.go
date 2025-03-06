package binary_search

import (
	"testing"
)

type TestTimeMapTestCase struct {
	test func() (string, string, bool)
}

func TestTimeMap(t *testing.T) {
	testCases := []TestTimeMapTestCase{
		{
			func() (string, string, bool) {
				timeMap := NewTimeMap()

				timeMap.Set("alice", "happy", 1)
				val1 := timeMap.Get("alice", 1)

				if val1 != "happy" {
					t.Errorf("Testing TimeMap, failed at case %d, expecting %s, got %s", 1, "happy", val1)

					return "happy", val1, false
				}

				val2 := timeMap.Get("alice", 2)

				if val2 != "happy" {
					t.Errorf("Testing TimeMap, failed at case %d, expecting %s, got %s", 2, "happy", val2)

					return "happy", val2, false
				}

				timeMap.Set("alice", "sad", 3)
				val3 := timeMap.Get("alice", 3)

				if val3 != "sad" {
					t.Errorf("Testing TimeMap, failed at case %d, expecting %s, got %s", 3, "sad", val3)

					return "sad", val3, false
				}

				val4 := timeMap.Get("david", 3)
				if val4 != "" {
					t.Errorf("Testing TimeMap, failed at case %d, expecting %s, got %s", 4, "", val4)

					return "", val4, false
				}

				return "", "", true
			},
		},
	}

	for _, testCase := range testCases {
		testCase.test()
	}
}
