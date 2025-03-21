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

# Structs
- grouping data & functions into collections
- what are structs?
- creating and using structs
- adding methods to structs

In Go, structs are composite data types that allow you to group together different data fields under a single name. They are used for organizing data and can also encapsulate functions as methods. Here’s a breakdown of structs in Go:

## What are Structs?
- Definition: Structs are **user-defined types** that contain named fields (data members).
- Purpose: They allow you to **group together different types of data** into a single unit.

## Creating and Using Structs
To define a struct in Go, you use the type and struct keywords followed by the name of the struct and its fields:

```bash
type Person struct {
    firstName string
    lastName  string
    age       int
}
```

- Creating Instances: You can create instances (objects) of a struct using the struct literal syntax:

```bash
p := Person{
    firstName: "John",
    lastName:  "Doe",
    age:       30,
}
```

- Accessing Fields: Struct fields are accessed using dot notation (.)
```bash
fmt.Println("First Name:", p.firstName)
fmt.Println("Last Name:", p.lastName)
fmt.Println("Age:", p.age)
```

## Adding Methods to Structs
In Go, methods are functions associated with a type. You can define methods for structs by specifying the receiver type (the struct itself):

```bash
func (p Person) fullName() string {
    return p.firstName + " " + p.lastName
}
```

- **Receiver**: **(p Person) indicates that the fullName method is associated with the Person struct**.
- Usage: Methods are called using dot notation on instances of the struct:

## Example

```bash
package main

import "fmt"

type Person struct {
    firstName string
    lastName  string
    age       int
}

func (p Person) fullName() string {
    return p.firstName + " " + p.lastName
}

func main() {
    p := Person{
        firstName: "John",
        lastName:  "Doe",
        age:       30,
    }

    fmt.Println("First Name:", p.firstName)
    fmt.Println("Last Name:", p.lastName)
    fmt.Println("Age:", p.age)
    fmt.Println("Full Name:", p.fullName())
}
```

In this example:

- We define a Person struct with fields firstName, lastName, and age.
- We define a fullName method that concatenates the first and last names.
- We create an instance p of type Person and print its fields and the result of calling fullName().

Structs in Go are fundamental for organizing data and are often used in place of classes in other object-oriented languages like Java or Python. They provide a flexible way to manage and manipulate structured data within your programs.

# Struct Tags
In Go, struct tags are annotations attached to struct fields. They provide metadata that can be used by reflection to influence the behavior of libraries or frameworks, such as how data is marshaled to JSON or how values are mapped from databases. Struct tags are written as raw string literals immediately following the field declaration within a struct definition.

## Basic Example of Struct Tags
Here’s a simple example with struct tags used for JSON serialization:
```bash
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```
In this example:

The struct tags (json:"id", json:"name", and json:"age") tell the encoding/json package how to serialize and deserialize the struct fields when converting to/from JSON.

## Struct Tag Syntax
Struct tags are raw string literals in backticks and typically consist of key-value pairs separated by spaces. The key identifies the user of the tag (like json in the previous example), and the value specifies additional instructions (e.g., how the field should be represented).

> FieldName Type `key1:"value1" key2:"value2"`

You can have multiple key-value pairs in a single tag, separated by spaces:
```bash
type Person struct {
    FirstName string `json:"first_name" db:"first_name" validate:"required"`
}
```

In this example:

- json:"first_name" tells the JSON marshaller to use "first_name" as the key in JSON output.
- db:"first_name" could be used by an ORM to map this field to a database column called "first_name".
- validate:"required" could be used by a validation library to ensure that this field is present.

## Common Uses of Struct Tags
1. JSON/Marshalling Tags (json): These tags are used when converting structs to and from JSON format with the encoding/json package.

Example:
```bash
type Product struct {
    Name  string `json:"name"`
    Price int    `json:"price,omitempty"`
}
```
- omitempty: When converting to JSON, if Price is set to the zero value (0 for integers, empty string for strings, etc.), it will be omitted from the JSON output.

2. Database Tags (db): Struct tags are often used by ORMs to map struct fields to database columns.

Example:
```bash
type Book struct {
    Title  string `db:"title"`
    Author string `db:"author_name"`
    Pages  int    `db:"pages"`
}
```
In this example, the db tags tell the ORM what the corresponding column names are in the database.

3. Validation Tags (validate): Libraries like go-playground/validator use struct tags to define validation rules for struct fields.
```bash
type User struct {
    Email string `validate:"required,email"`
    Age   int    `validate:"gte=18"`
}
```

- required: The field must be present.
- email: The field must contain a valid email address.
- gte=18: The Age must be greater than or equal to 18.

4. Custom Struct Tags: You can define your own struct tags if you write a library that uses reflection to inspect struct fields.
Example:
```bash
type Config struct {
    Host string `config:"server_host"`
    Port int    `config:"server_port"`
}
```
You can use the Go reflect package to access and use these tags in your own logic.

## Accessing Struct Tags
To access struct tags programmatically, you can use the reflect package. Here’s an example of how to retrieve the json struct tag for a field:

```bash
package main

import (
    "fmt"
    "reflect"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    u := User{}
    t := reflect.TypeOf(u)
    field, _ := t.FieldByName("Name")
    
    // Access the "json" tag
    jsonTag := field.Tag.Get("json")
    fmt.Println("JSON Tag for 'Name':", jsonTag)
}
```

This will output:
> JSON Tag for 'Name': name

## Tag Variants
1. Omitting Fields from Tags: You can omit a struct field from being serialized by using a hyphen (-) in a tag:
```bash
type User struct {
    Password string `json:"-"`
}
```

2. Customizing Tags for Special Needs: You can use special options in struct tags. For example, when using gorm (an ORM library), tags might look like:
```bash
type Employee struct {
    ID   int    `gorm:"primary_key"`
    Name string `gorm:"type:varchar(100)"`
}
```
- primary_key: Marks the ID field as the primary key.
- type:varchar(100): Specifies the SQL type for the Name field.

## Struct Tag Rules
1. Tags should be quoted as string literals in backticks (`).
2. Each key can only appear once in a struct tag. If a key is repeated, only the last one will be used.
3. Tags are optional. If you don’t need them, you can define structs without tags.

## Conclusion
Struct tags in Go are a powerful way to add metadata to struct fields, allowing you to control how data is marshaled, validated, or mapped to other systems. They are often used by libraries (e.g., json, database/sql, validator) but can also be custom-defined for specific use cases. Understanding struct tags and how to use them with reflection gives you flexibility when working with data in Go.

# Concurrency
**Concurrency** in *Go* refers to the ability to run multiple tasks independently but not necessarily at the same time. It's a way to efficiently manage multiple operations using fewer resources.

Go provides built-in support for concurrency using goroutines and channels.

## 🧑‍💻 Key Concepts in Go Concurrency
1. Goroutines
- Lightweight threads managed by the Go runtime.
- Execute functions concurrently.
- Created using the go keyword.

2. Channels
- Used to communicate between goroutines safely.
- Provide a way to send and receive data.

3. Wait Groups
- Synchronize multiple goroutines.
- Mutex
- Prevent race conditions using a lock.

## ✅ 1. Goroutines
A goroutine is a function that runs concurrently. It’s like a lightweight thread.

Example:
```go
package main
import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printMessage("Hello from Goroutine!") // Run concurrently
	printMessage("Hello from Main!")
}
```

## ✅ 2. Channels
Channels are used to safely share data between goroutines.

- Unbuffered Channels: Block until both sender and receiver are ready.
- Buffered Channels: Have a specified capacity and won't block until full.

Example:
```go
package main
import "fmt"

func square(number int, ch chan int) {
	result := number * number
	ch <- result // Send result to channel
}

func main() {
	ch := make(chan int)
	go square(4, ch) // Run square in a Goroutine

	fmt.Println("Square is:", <-ch) // Receive result from channel
}
```

## ✅ 3. Wait Groups
Use Wait Groups to wait for multiple goroutines to complete before moving on.

Example:
```go
package main
import (
	"fmt"
	"sync"
)

func printTask(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark as done
	fmt.Printf("Task %d is completed\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Add to the wait group
		go printTask(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks completed")
}
```

## ✅ 4. Mutex for Synchronization
A Mutex (short for mutual exclusion) ensures that only one goroutine accesses a shared resource at a time.

Example:
```go
package main
import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
```

## ✅ Concurrency vs Parallelism
- Concurrency: Tasks start, pause, and resume to make progress. Tasks may not run at the same time.
- Parallelism: Tasks run at the exact same time using multiple CPU cores.

Go schedules goroutines using GOMAXPROCS (number of CPU cores). By default, it uses all available cores.

## ✅ Best Practices for Concurrency in Go
- Use channels to safely share data between goroutines.
- Use wait groups to wait for multiple goroutines to finish.
- Use mutex to protect shared resources.
- Avoid data races by using proper synchronization mechanisms.
- Limit the number of goroutines to prevent resource exhaustion.

### 🔎 Analogy to Remember
1. Concurrency is like a chef cooking multiple dishes:
- The chef prepares multiple dishes by switching between them (checking the oven, stirring a pot).
2. Parallelism is like multiple chefs cooking at the same time:
- Each chef is responsible for one dish, cooking simultaneously.