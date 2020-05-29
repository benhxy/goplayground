* for _, t := range ts reuses the same address for t. So goroutine or defer statement might behave unexpected.
* Channel can be used as queue!
* We need a way to tell an unknown and unbounded number of goroutines to stop sending their values downstream. In Go, we can do this by closing a channel, because a receive operation on a closed channel can always proceed immediately, yielding the element type's zero value.
* Go string comparison is by value
* string is a value type, so cannot be null
* zero value for string is ""
* Go does have an extensive library, called the runtime, that is part of every Go program. The runtime library implements garbage collection, concurrency, stack management, and other critical features of the Go language.
* if a function is returned before the sub goroutine is executed, the goroutine will terminates and can't finish execution
* Passing slice pointer as argument is always preferred as any changes in underlying array can be reflected, including reallocation
* methods with pointer receivers take either a value or a pointer as the receiver when they are called
* Arguments are passed by copied value. To pass by reference, pass in a pointer
* initiate map: make(map[keyType]valueType)
* When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
* To create a set, make a map with bool value
* To return a struct pointer, return the address of the struct (&variable)
* Null check: (variable == nil)
* Adding parameter in front of function name makes it a method (similar to extension method)
* list is LinkedList, slice is ArrayList
* Named return values in a function are initialized with default value
* To return a struct, return type is the pointer of the struct
* To cast an element to a struct, do element.(*Struct)
* When assining a nil struct to an interface, because interface is a wrapper of {type, value}, the interface will not be nil.
* Nil slice (var s []int) is a slice with nil backing array. Empty slice (make([]int, 0)) is a slice pointing to an array with zero length. They are different.
* Nil map, channel, function are just a nil pointer.
* The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.