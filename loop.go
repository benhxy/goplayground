package main

import "fmt"

type testStruct struct {
	ID int
}

func printId(t *testStruct) {
	fmt.Println(t.ID)
}

func loopWithStruct() {
	ts := []testStruct{{1}, {2}, {3}}
	for _, t := range ts {
		fmt.Printf("%p", &t)
		fmt.Println("")
		defer printId(&t)
	}
}

func loopWithPointer() {
	ts := []*testStruct{&testStruct{1}, &testStruct{2}, &testStruct{3}}
	for _, t := range ts {
		fmt.Printf("%p", &t)
		fmt.Println("")
		defer printId(t)
	}
}

func loopTest() {
	fmt.Println("loopWithStruct()")
	loopWithStruct()
	fmt.Println()
	fmt.Println("loopWithPointer()")
	loopWithPointer()
}
