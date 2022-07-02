# EXERCISES

## [Calculations](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/artihmetic-operations)

```go
// ---------------------------------------------------------
// EXERCISE: Do Some Calculations
//
//  1. Print the sum of 50 and 25
//  2. Print the difference of 50 and 15.5
//  3. Print the product of 50 and 0.5
//  4. Print the quotient of 50 and 0.5
//  5. Print the remainder of 25 and 3
//  6. Print the negation of `5 + 2`
//
// EXPECTED OUTPUT
//  75
//  34.5
//  25
//  100
//  1
//  -7
// ---------------------------------------------------------
```


## [Comment out](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/comment-out)

```go
// ---------------------------------------------------------
// EXERCISE: Comment out
//
//  Use single and multiline comments to comment Printlns.
//
// EXPECTED OUTPUT
//  You shouldn't see any output after you're done.
// ---------------------------------------------------------
```


## [Rename Imports](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/importing)

```go
// ---------------------------------------------------------
// EXERCISE: Rename imports
//
//  1- Import fmt package three times with different names
//
//  2- Print a few messages using those imports
//
// EXPECTED OUTPUT
//  hello
//  hey
//  hi
// ---------------------------------------------------------
```


## [Naked Expression](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/naked-expression)

```go
// ---------------------------------------------------------
// EXERCISE: Naked Expression
//
//  1. Try to type just "Hello" on a line.
//  2. Do not use Println
//  3. Observe the error
//
// ---------------------------------------------------------
```


## [Operators combine the expressions](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/operators-combine)

```go
// ---------------------------------------------------------
// EXERCISE: Operators combine the expressions
//
//  Print the expected output below using the string
//  concatenation operator.
//
// HINT
//  Use + operator multiple times to create "Hello!!!?".
//
// EXPECTED OUTPUT
//  "Hello!!!?"
// ---------------------------------------------------------
```


## [Packages](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/packages)

```go
// ---------------------------------------------------------
// EXERCISE: Use your own package
//
//  Create a few Go files and call their functions from
//  the main function.
//
//  1- Create main.go, greet.go and bye.go files
//  2- In main.go: Call greet and bye functions.
//  3- Run `main.go`
//
// HINT
//  greet function should be in greet.go
//  bye function should be in bye.go
//
// EXPECTED OUTPUT
//  hi there
//  goodbye
// ---------------------------------------------------------
```


## [Parse Arg Numbers](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/parse-arg-numbers)

```go
// ---------------------------------------------------------
// EXERCISE: Parse Arg Numbers
//
//  Use strconv.ParseInt function to get int8, int16, and
//  int32, and int64 values from command-line.
//
// HINT
//  The third argument to ParseInt function represents
//  the bitsize.
//
//  So, giving it 8 returns an int8 convertable value;
//  whereas 16 returns an int16 convertable value.
//
//  Please explore the documentation of ParseInt function
//  and learn how it works.
//
// EXPECTED OUTPUT
//   When runned like this:
//     go run main.go 50 25000 2000000 50000000000000000 00000100
//
//   It should return this:
//     int8 value is : 50
//     int16 value is: 25000
//     int32 value is: 2000000
//     int64 value is: 50000000000000000
//     00000100 is: 4
// ---------------------------------------------------------
```


## [Print Go Version](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/print-goversion)

```go
// ---------------------------------------------------------
// EXERCISE: Print the Go Version
//
//  1. Look at the runtime package documentation
//  2. Find the func that returns the Go version
//  3. Print the Go version by calling that func
//
// HINT
//  It's here: https://golang.org/pkg/runtime
//
// EXPECTED OUTPUT
//  "go1.10"
// ---------------------------------------------------------
```


## [Print Hexes](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/print-hexes)

```go
// ---------------------------------------------------------
// EXERCISE: Print hexes
//
//  1. Print 0 to 9 in hexadecimal
//
//  2. Print 10 to 15 in hexadecimal
//
//  3. Print 17 in hexadecimal
//
//  4. Print 25 in hexadecimal
//
//  5. Print 50 in hexadecimal
//
//  6. Print 100 in hexadecimal
//
// EXPECTED OUTPUT
//  0 1 2 3 4 5 6 7 8 9
//  10 11 12 13 14 15
//  17
//  25
//  50
//  100
//
// NOTES
//  https://stackoverflow.com/questions/910309/how-to-turn-hexadecimal-into-decimal-using-brain
//
// https://simple.wikipedia.org/wiki/Hexadecimal_numeral_system
//
// ---------------------------------------------------------
```


## [Print Names](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/print-names)

```go
// ---------------------------------------------------------
// EXERCISE: Print names
//
//  Print your name and your best friend's name using
//  Println twice
//
// EXPECTED OUTPUT
//  YourName
//  YourBestFriendName
//
// BONUS
//  Use `go run` first.
//  And after that use `go build` and run your program.
// ---------------------------------------------------------
```


## [Print Literals](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/print-the-literals)

```go
// ---------------------------------------------------------
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool constants
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------
```

## [Scopes](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/scopes)-[Scopes2](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/scopes-2)

```go
// ---------------------------------------------------------
// EXERCISE: Try the scopes
//
//  1. Create two files: main.go and printer.go
//
//  2. In printer.go:
//     1. Create a function named hello
//     2. Inside the hello function, print your name
//
//  3. In main.go:
//     1. Create the usual func main
//     2. Call your function just by using its name: hello
//     3. Create a function named bye
//     4. Inside the bye function, print "bye bye"
//
//  4. In printer.go:
//     1. Call the bye function from
//        inside the hello function
// ---------------------------------------------------------
```

## [Shy Semicolon](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/shy-semicolon)

```go
// ---------------------------------------------------------
// EXERCISE: Shy Semicolons
//
//  1. Try to type your statements by separating them using
//     semicolons
//
//  2. Observe how Go fixes them for you
//
// ---------------------------------------------------------
```


## [Type Conversion](https://github.com/Property-Finder-Patika/week-2-homework-1-emretask1n/tree/main/Exercises/type-conversion)

```go
// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//
//  Fix the codes by using the conversion expression.
//
// EXPECTED OUTPUT-1
//  15.5

// EXPECTED OUTPUT-2
//  10.5

// EXPECTED OUTPUT-3
//  5.5

// EXPECTED OUTPUT-4
//  9.5

// EXPECTED OUTPUT-5
//  1127
// ---------------------------------------------------------
```