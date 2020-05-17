package landmine

import "fmt"

// This is to demonstrate the behavior of i,e := range []element when element's type is struct.

type testStruct struct {
	ID int
}

func printId(t *testStruct) {
	fmt.Printf("Address=%p, Value=%v", &t, t.ID)
	fmt.Println()
}

func loopStructElementsByRange() {
	ts := []testStruct{{1}, {2}, {3}}

	// When elements are assigned from a range,
	// they are assigned to the same variable (memory address),
	// causing a de facto closure into printId().
	for _, t := range ts {
		fmt.Printf("ParamAddress=%p, ParamValue=&v", &t, t.ID)
		fmt.Println()
		defer printId(&t)
	}
}

func loopStructElementsByIndex() {
	ts := []testStruct{{1}, {2}, {3}}

	// One solution is to assign a local variable inside the loop.
	// This forces a new memory assignment.
	for i := range ts {
		t := ts[i]
		fmt.Printf("ParamAddress=%p, ParamValue=&v", &t, t.ID)
		fmt.Println()
		defer printId(&t)
	}
}

func loopPointerElementsByRange() {
	ts := []*testStruct{&testStruct{1}, &testStruct{2}, &testStruct{3}}

	// Passing the pointer does not have this issue,
	// as the pointer value is copied into the function by parameter.
	for _, t := range ts {
		fmt.Printf("ParamAddress=%p, ParamValue=&v", t, t.ID)
		fmt.Println()
		defer printId(t)
	}
}

func loopTest() {
	fmt.Println("loopStructElementsByRange()")
	loopStructElementsByRange()
	fmt.Println()

	fmt.Println("loopStructElementsByIndex()")
	loopStructElementsByIndex()
	fmt.Println()

	fmt.Println("loopPointerElementsByRange()")
	loopPointerElementsByRange()
}
