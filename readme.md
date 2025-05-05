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
    - error handling
    - panic & recovery
    - modules & packages
    - structs
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
        - Cheap (~2KB) vs OS Thread (~4MB)
    - Concurrency support is built in the language
        - "go" keyword
        - "chan" datatype
        - "<-" operator
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