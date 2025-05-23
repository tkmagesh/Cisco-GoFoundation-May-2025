# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
| What | When |
| ----- | ---- |
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins) |
| Lunch Break | 12:30 PM (1 hour) |
| Tea Break | 3:00 PM (20 mins) |
| Wind up | 5:00 PM |

## Methodology
- NO powerpoint
- Code & Discussion
- Floor is always open for Questions

## Software Requirements
- Go Tools (https://go.dev/dl)
    ```shell
    go version
    ```
- Visual Studio Code (https://code.visualstudio.com)

## Repository
- https://github.com/tkmagesh/cisco-gofoundation-may-2025

## Daywise Breakup
### Day-01
    - Data Types
    - Language Constructs 
        - variables
        - constants
        - if else, select case
        - for
        - functions
### Day-02
    - collections
    - error handling
    - panic & recovery
    - modules & packages
    - structs
        - struct composition
### Day-03
    - methods
    - type assertion
    - abstraction (interfaces)
    - concurrency
    
## Why Go?
- Simplicity
    - ONLY 25 keywords
    - Only one way doing things
        - var, :=
        - if else, switch case
        - for
        - function
        - type
        - package
    - No access modifiers (NO private/public/protected keywords)
    - No reference types (everything is a value)
    - No pointer arithmatic supported by the language ( however, support offered through standard library "unsafe" package)
    - No classes (Only structs)
    - No class inheritance (Only struct composition)
    - No exceptions (Only errors & errors are just values)
    - No try..catch..finally
    - No implicit type conversion
- Performance
    - Comparable with C++
    - Close to hardware
        - Build for specific platform
        - Compiler has native support for cross compilation
- Managed Concurrency (**TBD**)
    - Builtin Scheduler to manage concurrent operations
    - "Goroutine" for concurrent operations
        - Cheap (~2KB) vs OS Thread (~2MB)
    - Concurrency support is built in the language
        - "go" keyword
        - "chan" datatype
        - "<-" channel operator
        - "range" construct
        - "select-case" construct
    - Standard library support
        - "sync" package
        - "sync/atomic" package
![image](./images/WhyGo.png)

## Compilation
```shell
go build [filename.go]

# create a binary in a different name
go build -o [binary_name] [filename.go]
```

## Compile & Execute
```shell
go run [filename.go]
```

## Cross Compilation
### list all the go compiler's env variables
```shell
go env
```

### list specific env variables
```shell
go env [var_1] [var_2]
# ex
go env GOPATH GOBIN
```

### Change env variables
```shell
go env -w [var_1]=[value_1] [val_2]=[value_2] ...
```

### Env variables for cross compilation
- GOOS
- GOARCH

### To get the list of supported platforms (for cross compilation)
```shell
go tool dist list
```

### Cross compilation example
```shell
GOOS=windows GOARCH=amd64 go build [filename.go]
```

## Data Types
- string
- bool
- integers
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points
    - float32
    - float64
- complex
    - complex64 ( real[float32] + imaginary[float32] )
    - complex128 ( real[float64] + imaginary[float64] )
- alias
    - byte (alias for unsigned int)
    - rune (alias for unicode code point)

### Zero values

| Data Type | Zero value |
------------ | ------------- |
|int family     | 0 |
|uint family    | 0 |
|float family    | 0 |
|complex family | (0+0i) |
|string         | "" (empty string) |
|bool           | false |
|byte           | 0 |
|interface      | nil |
|pointer        | nil |
|function       | nil |
|struct         | struct instance |



## Standard Library Packages
- Documentation
    - https://pkg.go.dev/std
- Source Code
    - $(go env GOROOT)/src 

## Package Vs Function scope
### Package scope
- CANNOT use ":=" (Only "var" is allowed)
- CAN have unused variables
### Function scope
- CAN use ":="
- CANNOT have unused variables

## Constants
- CAN have unused constants (in both "function" and "package" scope)

## iota
- Constants group with auto-generated values

## Constructs
- if else
- switch case
- for

## Functions
- Named results
- Variadic function
- Anonymous function
    - no names
    - have to be immediately invoked
- Higher Order functions
    - Assign a function as a value to a variable
    - Pass a function as an argument
    - Return a function as a return value
- Deferred functions
    - postpone the execution of a function until the current function execution is completed

## Collection Types
### Array
- Fixed sized typed collection

### Slice
- Varying sized typed collection
- append()
- len()
- cap()
- make()
![image](./images/slices.png)

### Map
- Typed collection of key/value pairs
- Has to be initialized with make() 
- delete()

## Error Handling
- An Error is just a value "returned" from a function
- By convention, an error value should implement "error" interface
    - any object with "Error()" method
- Ways to create an error
    - errors.New()
    - fmt.Errorf()
    - Custom type implementing "error" interface

## Panic & Recovery
### Panic
- State of the application where the application execution is unable to proceed further.
- When a "panic" occurs, the application is shutdown after executing all the  functions (deferred) that are already deferred
- Use "panic()" to programmatically raise a panic
    - a panic is typically raised with an error (that led to the panic scenario)

### Recovery
- The "recover()" function can be used to recover from a panic
- The "recover()" function returns the error that resulted in the panic
- Apt to use the "recover()" function in "deferred" functions

## Modules & Packages
### Module
- Any code that has to be versioned and deployed together
- A module is typically a folder with "go.mod" file
- **go.mod** file (manifest file) contains:
    - the name of the module
    - the go version targetted
    - references to other dependencies


#### Create a `go.mod` file
```shell
go mod init <module_name>
```
- by convention, the module name should include the repo path (not mandatory)

#### Compile & Execute a module
```shell
go run .
```

#### Compile a module
```shell
go build .
# OR
go build -o [binary-name] .
```


### Package
- internal code organization of a module
- typically folders
- can also be nested
- All the code in a package (across files) are considered to be belonging to the "package"
- Accessibility is determined by the naming convention
    - Public - the entity name must start with **uppercase**
    - Private - the entity name must start with **lowercase**

### Using Open Source
#### Repo
- https://pkg.go.dev
- https://github.com/avelino/awesome-go

#### Installation

[downloaded in the GOPATH/pkg folder]

```shell
# Execute the command in the go.mod folder so that the go.mod file can be updated
go get <module-name>
# ex
go get github.com/fatih/color
```

#### Update references in the go.mod file 
```shell
go mod tidy
```

#### To download the dependencies referenced in the `go.mod` file
```shell
go mod download
```

#### Other useful module commands
```shell
go mod graph
go mod why <module-name>
go list -m all
```

#### Go module command reference
- https://go.dev/ref/mod

## OO Programming in Go
### Struct
- Used to encapsulate related data & behavior together
### Abstractions
#### Interfaces
- Interfaces are "implicitly" implemented in Go
- Interfaces can be composed

## Concurrency
### Concurrent Design
- The application is designed in such a way that it has more than one execution path
### Concurrency in Go
- Independent Execution path is represented as a "goroutine" 
- goroutines are cheap (2 KB)
- Application build is embedded with a "scheduler"
- Builtin scheduler schedules the goroutines through the OS threads
- Scheduler follows "cooperative multitasking"

## Cooperative Multitasking
![image](./images/cooperative-multitasking.png)

## Pre-emptive Multitasking
![image](./images/preemptive-multitasking.png)

## Concurrency vs Parallelism
![image](./images/concurrency-vs-parallelism.png)

## Go Concurrency Model
![image](./images/go-concurrency.png)

## Threads Vs Goroutines
![image](./images/threads-vs-goroutines.png)

## Goroutine synchronization using `WaitGroup`
- `WaitGroup` is a semaphore based counter
- Has the ability to block the execution of a function until the counter becomes `0`

## Data Race
### To detect data race
```shell
go run -race <filename.go | .>
```
```shell
go build -race <filename.go | . >
```

## Communication between goroutines
### Channel
- datatype built to enable communication between goroutines
#### Declaration
```go
var ch chan int
```
#### Initialization
```go
ch = make(chan int)
```
#### Send Operation
```go
ch <- 100
```
#### Receive Operation
```go
data := <- ch
```


