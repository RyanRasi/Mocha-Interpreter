# Mocha-Interpreter

**An interpreter language written in go**

Mocha is a programming language that is implmented with Go.

Currently Work In Progress.

Mocha has/is planned to have the following features:

- [x] C-like syntax
- [x] Variable bindings
- [x] Integers and booleans
- [x] Arithmetic expressions
- [x] Built-in functions
- [x] First-class and higher-order functions
- [x] Closures
- [x] A string data structure
- [x] An array data structure
- [x] A hash data structure
- [ ] Attach line number, column number and filename to a token
- [ ] Support full unicode range
- [ ] Allow ! in identifiers and keywords
- [ ] Support Floats / Hex Notation / Octal Notation
- [ ] Comparators e.g. AND OR NOT
- [ ] Add instant web server
- [x] Add .mocha file extension support
- [x] If statements
- [x] Else statements
- [ ] Elif statements
- [x] Arrays
- [x] Array Literals
- [ ] For loop
- [ ] Do while loop
- [ ] Switch case statements
- [ ] Postfix operators (e.g. ++)
- [x] Built-in library
- [ ] True integer division support
- [x] Hash Map/Dictionary
- [x] Automatic formatting of custom user files on run

To run the IDE do...

```
go run main.go
```

Else to load a programming file do...

```
go run main.go test.mocha
```

Replace test.mocha with your created file and place it within the src folder

#Built In Functions

Output to console

```
consoleOut("String or variable");
```

Arrays

```
first(); - Get the first element of array
last(); - Get the last element of array
rest(); - Get all elements of an array excluding the first element
push('String or Integer', Array); - Append an element to the end of an array
sort(); - Sorts an array in decending order
```

Math Statements

```
add(); - Add multiple integers together
```

help(); - Displays help commands

```
exit(); - Exits the program safely
```
