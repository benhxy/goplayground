package landmine

import (
	"fmt"
	"testing"
)

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
		fmt.Printf("ParamAddress=%p, ParamValue=%v \n", &t, t.ID)
		defer printId(&t)
	}
}

func loopStructElementsByIndex() {
	ts := []testStruct{{1}, {2}, {3}}

	// One solution is to assign a local variable inside the loop.
	// This forces a new memory assignment.
	for i := range ts {
		t := ts[i]
		fmt.Printf("ParamAddress=%p, ParamValue=%v \n", &t, t.ID)
		defer printId(&t)
	}
}

func loopPointerElementsByRange() {
	ts := []*testStruct{&testStruct{1}, &testStruct{2}, &testStruct{3}}

	// Passing the pointer does not have this issue,
	// as the pointer value is copied into the function by parameter.
	for _, t := range ts {
		fmt.Printf("ParamAddress=%p, ParamValue=%v \n", t, t.ID)
		defer printId(t)
	}
}

func TestLoopOverRange(t *testing.T) {
	fmt.Println("loopStructElementsByRange()")
	loopStructElementsByRange()

	fmt.Println("loopStructElementsByIndex()")
	loopStructElementsByIndex()

	fmt.Println("loopPointerElementsByRange()")
	loopPointerElementsByRange()

	/*
		loopStructElementsByRange()
		ParamAddress=0xc000014128, ParamValue=1
		ParamAddress=0xc000014128, ParamValue=2
		ParamAddress=0xc000014128, ParamValue=3
		Address=0xc000006038, Value=3
		Address=0xc000006040, Value=3
		Address=0xc000006048, Value=3

		loopStructElementsByIndex()
		ParamAddress=0xc000014188, ParamValue=1
		ParamAddress=0xc000014198, ParamValue=2
		ParamAddress=0xc0000141a8, ParamValue=3
		Address=0xc000006050, Value=3
		Address=0xc000006058, Value=2
		Address=0xc000006060, Value=1

		loopPointerElementsByRange()
		ParamAddress=0xc0000141e8, ParamValue=1
		ParamAddress=0xc0000141f0, ParamValue=2
		ParamAddress=0xc0000141f8, ParamValue=3
		Address=0xc000006068, Value=3
		Address=0xc000006070, Value=2
		Address=0xc000006078, Value=1
	*/
}
