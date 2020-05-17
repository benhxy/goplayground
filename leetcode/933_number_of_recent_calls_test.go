package goleetcode

import "testing"

func TestAllCurrentPings(t *testing.T) {
	recentCounter := RecentCounterConstructor()

	result1 := recentCounter.Ping(0)
	if result1 != 1 {
		t.Errorf("Expected %v, actual is %v", 1, result1)
	}

	result2 := recentCounter.Ping(1000)
	if result2 != 2 {
		t.Errorf("Expected %v, actual is %v", 1, result1)
	}

	result3 := recentCounter.Ping(4000)
	if result3 != 2 {
		t.Errorf("Expected %v, actual is %v", 1, result1)
	}
}
