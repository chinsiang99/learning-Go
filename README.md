# learning-Go
this repo is for me to learn Go

# What is Go
Go is an open-source programming language developed by Google

## Popular Use-cases
- newtworking & APIs
- microservices
- CLI tools

# Go Essentials
- understanding the key components of a Go program
- working with values & types
- creating & executing functions
- controlling execution with control structures

# Go Types & Null Values

## Basic Types
Go comes with a couple of built-in basic types:

1. int => A number WITHOUT decimal places (e.g., -5, 10, 12 etc)

2. float64 => A number WITH decimal places (e.g., -5.2, 10.123, 12.9 etc)

3. string => A text value (created via double quotes or backticks: "Hello World", `Hi everyone')

4. bool => true or false

But there also are some noteworthy "niche" basic types which you'll typically not need that often but which you should still know about:

1. uint => An unsigned integer which means a strictly non-negative integer (e.g., 0, 10, 255 etc)

2. int32 => A 32-bit signed integer, which is an integer with a specific range from -2,147,483,648 to 2,147,483,647 (e.g., -1234567890, 1234567890)

3. rune => An alias for int32, represents a Unicode code point (i.e., a single character), and is used when dealing with Unicode characters (e.g., 'a', 'ñ', '世')

4. uint32 => A 32-bit unsigned integer, an integer that can represent values from 0 to 4,294,967,295

5. int64 => A 64-bit signed integer, an integer that can represent a larger range of values, specifically from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

There also are more types like int8 or uint8 which work in the same way (=> integer with smaller number range)

## Null Values
All Go value types come with a so-called "null value" which is the value stored in a variable if no other value is explicitly set.

For example, the following int variable would have a default value of 0 (because 0 is the null value of int, int32 etc):

var age int // age is 0
Here's a list of the null values for the different types:

1. int => 0

2. float64 => 0.0

3. string => "" (i.e., an empty string)

4. bool => false

# fmt.Scan() Limitations

The fmt.Scan() function is a great function for easily fetching & using user input through the command line.

But this function also has an important limitation: You can't (easily) fetch multi-word input values. Fetching text that consists of more than a single word is tricky with this function.

For the moment, we only need single words / digits as input, so that's no problem.

Later in the course, when we work on projects where more complex input values are required, you'll therefore learn about an alternative to fmt.Scan().

# Conditional For Loops
Besides the for variations introduced in the last lectures, there also is another common variation (which will also be shown again later in the course):

```bash
for someCondition {
  // do something ...
}
```

someCondition is an expression that yields a boolean value or a variable that contains a boolean value (i.e., true or false).

In that case, the loop will continue to execute the code inside the loop body until the condition / variable yields false.

# Working with Go Packages
- splitting code across **multiple files**
- splitting files across **multiple packages**
- **importing & using** custom packages

# Understanding Pointers
- working with addresses instead of values
- what are pointers?
- why does this feature exist?
- how can you work with pointers?

## What Are Pointers?
In Go, a pointer is a variable that holds the memory address of another variable. Rather than containing a value directly, a pointer refers to the location in memory where that value is stored.

## Why Do Pointers Exist?
Pointers are a fundamental feature of Go (and many programming languages) for the following reasons:

1. Efficiency: Passing pointers allows large data structures to be passed around efficiently without copying.
2. Modifying Data: By passing a pointer, functions can modify the original variable's data.
3. Memory Management: Pointers are essential for managing dynamic memory allocation.
4. References: Pointers enable sharing references to data between different parts of a program.

## Working with Addresses Instead of Values
Typically, variables in Go store values directly. However, with pointers, you store the memory address of a variable. This allows you to access or modify the value indirectly.

For example, instead of directly holding an integer, a pointer to an integer stores the address where that integer resides in memory.

## How to Work with Pointers
### Declaring a Pointer
To declare a pointer in Go, you use the * symbol along with the type of the value the pointer will refer to. The pointer doesn't store a value but rather an address of a value.

> var ptr *int // declares a pointer to an int

### Dereferencing a Pointer
Dereferencing a pointer means accessing the value the pointer refers to. You can do this using the * operator.

```bash
var x int = 10
var ptr *int = &x  // ptr holds the address of x

fmt.Println(*ptr)  // dereferences ptr to print the value of x, which is 10
```

### Passing Pointers to Functions
When passing a variable to a function by pointer, the function can modify the original variable's value. This avoids making a copy of the data and allows in-place modifications.
```bash
func increment(ptr *int) {
    *ptr = *ptr + 1
}

func main() {
    x := 5
    increment(&x)   // passing the address of x
    fmt.Println(x)  // x is now 6
}
```

### Nil Pointers
A pointer that is not initialized points to nil. Dereferencing a nil pointer will cause a runtime panic, so it is essential to check for nil before dereferencing.

```bash
var ptr *int
if ptr != nil {
    fmt.Println(*ptr)
}
```

### basic pointer usage 1
```bash
package main

import "fmt"

func main() {
    x := 100
    ptr := &x  // pointer to x

    fmt.Println("Value of x:", x)          // prints 100
    fmt.Println("Address of x:", ptr)      // prints address of x
    fmt.Println("Value at ptr:", *ptr)     // dereferences ptr to get the value of x

    *ptr = 200  // modifies x through the pointer
    fmt.Println("Updated value of x:", x)  // prints 200
}
```

### basic pointer usage 2
```bash
package main

import "fmt"

func updateValue(val *int) {
    *val = 42  // modifying the original value
}

func main() {
    x := 10
    updateValue(&x)
    fmt.Println(x)  // prints 42
}
```

## Summary
Pointers in Go are a powerful tool for optimizing memory usage and enabling more flexible programming. By working with pointers, you can improve performance, manage data more effectively, and allow functions to modify variables in a safe, controlled manner.

# Pointers vs Values

In Go, whether to use pointers or values when passing variables to functions depends on the context and the specific use case. Here's a breakdown of when each might be more appropriate

## When to Use Pointers

1. Modifying the Original Value:

- If the function needs to modify the original value (not just work with a copy), you should pass a pointer.
- Example: Updating a struct's field or changing a variable's value.

```bash
func modify(val *int) {
    *val = 10 // modify the original value through the pointer
}

x := 5
modify(&x)
fmt.Println(x) // x is now 10
```

2. Avoiding Expensive Copies:

- If you're passing a large struct or data type, passing by value can be inefficient as it involves copying the entire object. Using pointers avoids this by passing the memory address instead.
- Example: Passing a large struct like []byte or a complex data type.

```bash
type LargeStruct struct {
    data [1000]int
}

func process(large *LargeStruct) {
    // process without copying the entire struct
}
```

3. Nil Values:

- If you need to represent the absence of a value (i.e., nil), you can only do this with pointers. Passing a nil pointer is a way to signify "no value."
- Example: Optional parameters.

```bash
func checkStatus(status *string) {
    if status == nil {
        fmt.Println("Status is nil")
    } else {
        fmt.Println(*status)
    }
}
```

4. Maintaining State Between Function Calls:

- Pointers allow state to be shared and updated across multiple function calls. Since the memory address is passed, any modifications in one function persist in others.
- Example: Passing a pointer to manage shared data.

## When to Use Values:

1. Small or Simple Data Types:

- For small data types like integers, booleans, floats, and small structs, it's usually more efficient to pass them by value since the cost of copying is minimal.
- Example: Basic types like int, float64, bool.

```bash
func increment(x int) int {
    return x + 1 // works with a copy of x
}
```

2. Immutability:

- If the function doesn't need to modify the original data or ensure that changes persist outside the function, passing by value provides a "snapshot" of the data.
- Example: A function where you only need to read the data without altering it.

```bash
func printValue(x int) {
    fmt.Println(x) // doesn't need to modify x
}
```

3. Security and Simplicity:

- Passing by value guarantees that the original data won't be accidentally modified, which can be useful when immutability is desired.
- Example: Avoiding side effects by ensuring the original value remains unchanged.

```bash
func safeModify(x int) int {
    return x + 10 // does not affect the original variable
}
```

## General Guidelines:
Use Pointers if:

1. The data is large and copying would be expensive (e.g., large structs or slices).
2. You need to modify the data in the function and have those changes reflected outside the function.
3. You need to represent optional values (nil pointers).

Use Values if:

1. The data is small and the cost of copying is negligible (e.g., basic types like int, float64).
2. You want to ensure that the original data remains unchanged.
3. You're working with immutable data.

## Conclusion:
- Use pointers when you need to modify the original data, avoid copying large structures, or handle optional/nil values.
- Use values for small, immutable data or when you want to ensure that the original data remains unchanged.

# A Pointer's Null Value
All values in Go have a so-called "Null Value" - i.e., the value that's set as a default if no value is assigned to a variable.

For example, the null value of an int variable is 0. Of a float64, it would be 0.0. Of a string, it's "".

For a pointer, it's nil - a special value built-into Go.

nil represents the absence of an address value - i.e., a pointer pointing at no address / no value in memory.