## Notes on Go

* for _, t := range ts reuses the same address for t. So goroutine or defer statement might behave unexpected due to closure.
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
* Type assertion to interface instead of struct is the way to do IOC in Go, and helps breaking dependency.

## Steps to learn
1. The language itself
   1. Basic types (primitive, struct, interface, func, slice, map, chan).
   1. Control structures (if, for, switch).
   1. func syntax (receiver, parameter, return values, error, closure).
   1. Concurrency (sync.Mutex, goroutine).
   1. Solve some Leetcode using the data structures and control structures.
1. Toolchain
   1. Project structure and layout (https://github.com/golang-standards/project-layout).
   1. Build process (go build, gofmt, go run). Go doesn't have virtual machine or C runtime. Everything is build into a native binary.
   1. Unit testing and benchmarking (go test, testing.T, testing.B).
   1. Go module and dependency management (go mod, go get, go install).
   1. Web server and profiling (net/http, net/http/pprof).
1. Libraries, frameworks, and web applications
   1. Take a look at source code of important standard libraries (errors, context, sync, io).
   1. Logging (logrus).
   1. App configuration (viper).
   1. gRPC and protobuf
   1. Compile-time dependency injection (github.com/google/wire).
   1. And ... well ... Kubernetes.


## Tooling
* Go Module and package/dependency management
* pprof

## Resources
* [Effective Go](https://golang.org/doc/effective_go.html): The official and most detailed walkthrough of language features and behaviors. It explains a lot of "why"s. Also goo reference for writing _idiomatic_ Go code.
* [Go by Example](https://gobyexample.com/): Similar to Effective Go. But shorter and more readable as code comments.
* [Go go-to guid](https://yourbasic.org/golang/): Each article explains a concept or behavior. Clear and concise writing.
* [Go standard library source code](https://github.com/golang/go): Go source code is very short and clean. It's very useful to read through some important packages (fmt, errors, context, strconv) and see the definitions and coding styles.
* [justforfunc](https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw): Ex-Googler explaining Go ecosystem (language + tooling). Always a pleasure to watch smart and hot bear coding ;)