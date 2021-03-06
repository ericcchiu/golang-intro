<!-- omit in toc -->

# Go Foundations

- [Go Foundations](#go-foundations)
  - [**1. Introduction**](#1-introduction)
    - [Use Cases:](#use-cases)
    - [Why Go?](#why-go)
  - [**2. Setting Up Golang**](#2-setting-up-golang)
    - [Learning Objectives:](#learning-objectives)
    - [**Installation**](#installation)
      - [Windows Installation](#windows-installation)
      - [Linux Installation](#linux-installation)
        - [Step One: Install Go Language](#step-one-install-go-language)
        - [Step 2: Set up Go environment](#step-2-set-up-go-environment)
  - [**3. Types**](#3-types)
    - [**Concepts**](#concepts)
    - [**Intro**](#intro)
    - [**Numbers**](#numbers)
      - [Integers](#integers)
      - [Float](#float)
      - [Using Number Literals](#using-number-literals)
      - [Rune and Byte](#rune-and-byte)
    - [**Strings**](#strings)
    - [**Boolean**](#boolean)
      - [Operators](#operators)
  - [4. **Go Variables**](#4-go-variables)
    - [**Learning Objective:**](#learning-objective)
    - [**Variable Concepts**](#variable-concepts)
    - [**Skills**](#skills)
    - [**Variables**](#variables)
    - [**Declaration**](#declaration)
    - [**Short Declaration**](#short-declaration)
    - [**Multiple Variable Values and Declarations**](#multiple-variable-values-and-declarations)
    - [**Constants**](#constants)
    - [**Lab 1 - Declare and use Variables**](#lab-1---declare-and-use-variables)
    - [**Scope**](#scope)
    - [**Type Definitions**](#type-definitions)
  - [5. **Control Structures**](#5-control-structures)
    - [**If Statements**](#if-statements)
    - [**Switch Statements**](#switch-statements)
      - [Default](#default)
      - [Short and More Complex Evaluations](#short-and-more-complex-evaluations)
      - [Switch with no Conditon](#switch-with-no-conditon)
    - [**Lab 2: Switch Statement**](#lab-2-switch-statement)
    - [**For - One Loop to Rule Them All!**](#for---one-loop-to-rule-them-all)
      - [Code that should probably be in a loop](#code-that-should-probably-be-in-a-loop)
      - [Basic iterative for loop\*](#basic-iterative-for-loop)
      - [While style loop](#while-style-loop)
      - [Continue, Break, Return and Labels](#continue-break-return-and-labels)
        - [Break & Return](#break--return)
      - [Continue](#continue)
      - [Labels](#labels)
      - [Using Labels with `continue` and `break`](#using-labels-with-continue-and-break)
      - [Goto](#goto)
    - [**Lab 3: Loops**](#lab-3-loops)
    - [Putting it All Together: If, Switch and Loops](#putting-it-all-together-if-switch-and-loops)
  - [6. **Array, Slices and Maps**](#6-array-slices-and-maps)
    - [**Theory**](#theory)
    - [**Objectives**](#objectives)
    - [**Arrays**](#arrays)
      - [**Indexes**](#indexes)
      - [**Assigning Values**](#assigning-values)
        - [Specific Values](#specific-values)
        - [Initiliazing With Values](#initiliazing-with-values)
      - [**Reading Values**](#reading-values)
        - [Determining the Size of an Array](#determining-the-size-of-an-array)
      - [Quick Quiz](#quick-quiz)
    - [**Slices**](#slices)
      - [**Creating a slice from an array**](#creating-a-slice-from-an-array)
      - [**Sizing empty slices: Introducing make()**](#sizing-empty-slices-introducing-make)
      - [**Slice Functions**](#slice-functions)
        - [**append()**](#append)
        - [**copy()**](#copy)
      - [**Further Reading on Arrays and Slices**](#further-reading-on-arrays-and-slices)
    - [**Lab 1: Arrays and Slices**](#lab-1-arrays-and-slices)
    - [**Maps**](#maps)

---

## **1. Introduction**

Go(Golang) was created by Google in November of 2009 originally to be an amazing procedural language for scalable cloud software. It was built to be simple and fast.

### Use Cases:

- System Admin scripts
- Web app that serves HTML pages
- File server
- Scripts that deploy code to environments

### Why Go?

- It's simple. The overarching reason behind every feature is readability.

- It compiles down to binaries so deploying is simple.

- It comes with only the basic structures and methods. Add packages only as needed. This allows the source code to remain light weight.

- It comes with an auto formatter, so spacing is not an argument on teams.
- Basic testing is built-in

- Concurrency is made easy for the developer. No thread management needed. The design of communication eliminates race conditions.

- It comes with all the things you need to make a server.

- Version 1 is completely backwards compativle, as promised.

- Interfaces allow for abstraction, easy testability (through Dependency Injection), and the Adapter Pattern.

- Consisten documentation found at [golang.org/doc](https://golang.org/doc/)

---

## **2. Setting Up Golang**

### Learning Objectives:

- Installation of Golang
- Steps to set up Golang
- Go Modes: Go Modules and GOPATH
- Parts of a Go file
- Package, Import and Declaration sections of a Go file
- Building and running a Go file

### **Installation**

#### Windows Installation

1. To install Golang on a Windows system, download and install the latest 64-bit Go set for Microsoft Windows OS: [https://golang.org/install/](https://golang.org/doc/install).
2. Follow the instructions on the Go installation program.
3. Run the Command Prompt on your computer by searching for "cmd". Open the command line and type: `"go version"`.
4. One example of the output of `go version` is:

```Linux
go version go1.14.4 windows/amd64
```

#### Linux Installation

##### Step One: Install Go Language

Upagrade to apply the latest security updates on Ubuntu.

```Linux
sudo apt-get update
sudo apt-get -y upgrade
```

You need to download the Go binary file. Check out [https://golang.org/install/](https://golang.org/doc/install) for the latest version of go.

Then run:

```linux
cd /tmp

wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
```

Download the archive and install it to the desired location on the system such as /usr/local

```linux
sudo tar -xvf go1.14.2.linux-amd64.tar.gz

sudo mv go /user/local
```

##### Step 2: Set up Go environment

Set up the Go language environment variable so that Go can be used outside the `$GOPATH`.

Add /user/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile(for a system-wide installation) or \$HOME/.profile:

```linux
export PATH=$PATH:/user/local/go/bin
```

1. Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source \$HOME/.profile.

---

## **3. Types**

### **Concepts**

- What is a built-in type
- The number types
- The string type
- Boolean types

### **Intro**

Types are how we identify a data type in Go. There are many different types in Go, such as `struct` and `interface`.

### **Numbers**

In Go there are two major number types: **Integers** and **Floating point** numbers.

#### Integers

Integers are non-decimal based numbers that can be represented using one of the following types:

- The `int` types can represent any positive or negative non-decimal number
- The `uint` types can only be positive non-decimal numbers

_Sized Types:_

- `int8`, `int16`, `int32`, and `int64`
- `uint8`, `uint16`, `uint32`, and `uint64`

The 8, 16, 32, and 64 by each type represents the maximum size of the type in bits. `int8` has a max size of 8 bits and `uint64` has a max size of 64 bits.

Examples:

- Given a 64 bit processors `int` will default to `int64`
- Given a 32 bit processors `int` will default to `int32`

#### Float

Floating point numbers are decimal numbers. Floating point numbers are represented by either `float32` or `float64`. Like `int`, the 32 and 64 represent the bit capacity of the floating point number.

There are also `complex64` and `complex128` types that fall under floating point that we will not cover here.

#### Using Number Literals

A literal is a value written as is, and not assigned to any variables.

It is very common to use the following operators with numbers when doing various calculations.

| Operator | Description       |
| :------: | :---------------- |
|    +     | addition          |
|    -     | subtraction       |
|    \*    | multiplication    |
|    /     | division          |
|    %     | remainder/Modulus |

Example of using number literals and operators

```Golang
package main

import "fmt"

func main() {
    fmt.Println(200 - 1 + 3 - 4)
    fmt.Println(-10 + 4)
    fmt.Println(1.2 + 3.4)
}
```

#### Rune and Byte

The `rune` and `byte` types are also technically numbers as well and are alias to types covered above. Their actual use will be discussed in following sections. For now its good to note that `rune` is the same as an `int32` in type. It is used to represent UTF-8 values of string literals.

Similarly `byte` is an alias for `unint8`. Byte is most commonly seen in when working with strings, which in Golang, are an array of bytes.

---

### **Strings**

String types are used to represent text. They are made up of a series of characters. More specifically, strings are made up of `bytes`, where there is typically one byte per character.

_**Creating Byte Literals**_

There are two ways to create string literals. A string literal can be created by surrounding the characters with double quites `"` or "back ticks" `.

_**Double Quotes**_

- Example: `"Get Going"`
- Will translate escape escape characters. For eample `\n` is translated to a newline.

_**Back Ticks**_

- **Example:** `Get Going`
- Will not translate escape characters.
- Handy when creating literals that contain double quotes, such as when working with JSON strings.

_**Concatenated Strings**_

You often need to combine one more or more strings at some point. The quickest, most common way to do this is to use the `+` operator to combine the two strings.

**Example:** `"Go on " + "with your bad self."`

There are other methods of concating strings that may be more efficient when working with large strings, or gorups of strings; however, using the `+` is perfectly fine in most cases. The other mehods will not be covered in this course, but if you are interested in some of the other methods [this is a good blog presenting](https://www.hermanschaaf.com/efficient-string-concatenation-in-go/) some benchmarking and techniques.

Alternate concatenation methods:

- Method 1: Naive Appending with `+=`

  This is the most obvious approach to the problem. Use the concatenate operator (+=) to append each segment to the string. The thing to know about this is that strings are immutable in Go - each time a string is assigned to a variable a new address is assigned in memory to represent the new value. This is different from languages like C++ and BASIC, where a string variable can be modified in place. In Go, each time you append to the end of a string, it must create a new string and copy the contents of both the existing string and the appended string into it. As the strings you are manipulating become large this process becomes increasingly slow. In fact, this method of naive appending is O(N^2).

- Method 2: strings.Join

  The Go strings standard library package provides a Join function that takes a slice of strings and concatenates them. In our case, this requires first building a slice of strings and then calling the function on the slice. Building a slice takes time if you don’t already have one, and we therefore count this as part of the time taken to perform the concatenation in our benchmarks. The final concatenation will be O(N), but building the slice is O(N) (average) and uses O(N) memory.

  Here is an example of using `strings.Join`:

  ```Golang
  package main

  import (
      "fmt"
      "strings"
  )

  func main() {
      s := []string{}
      for i := 0; i < 1000; i++ {
          s = append(s, "a")
      }
      fmt.Println(strings.Join(s, ""))
  }
  ```

- Method 3: bytes.Buffer

  The last method we’ll be evaluating is using the Buffer type from the bytes package. bytes.Buffer implements io.Writer, and can concatenate strings in O(N) time without requiring us to build a new slice.

  Here is an example of using `bytes.Buffer`:

  ```Golang
  package main

  import (
      "bytes"
      "fmt"
  )

  func main() {
      var buffer bytes.Buffer

      for i := 0; i < 1000; i++ {
          buffer.WriteString("a")
  }

  fmt.Println(buffer.String())
  }
  ```

- **Working with String Literals**

  ```Golang
  package main
  imprt "fmt"
  func main() {
      fmt.Println("Part one " + "joined with " + "Part 2")
      fmt.Println("I'm a string Literal!\n", `So am I!`)
      fmt.Println("Printing a JSON String: %s \n", `{"color": "red"}`)
  }

  ```

---

### **Boolean**

Boolean values represented by either `true` or `false`.

The `bool` type can define a variable that represents either of these states.

#### Operators

There are five operators used with boolean values in Go.

| Operator | Description |
| :------: | :---------- |
|   `&&`   | And         |
|   `||`   | Or          |
|   `!`    | Not         |
|   `==`   | Equals      |
|   `==`   | Equals      |
|   `!=`   | Not Equals  |

```Golang
package main

import "fmt"

func main() {
    fmt.Println("this" == "that") // false
    fmt.Println( 1 ++ 1) // true
    fmt.Println("this" == "this") // true
    fmt.Println("this" != "that") // true
    fmt.Println("this" == "this" && 1 == 1) // true
    fmt.Println("this" == "this" && 1 == 2) // false
    fmt.Println("this" == "this" || 1 == 2) // false
}
```

---

## 4. **Go Variables**

### **Learning Objective:**

### **Variable Concepts**

- Variable Declarations
- Short Variable Declarations
- Constants
- Multiple Variables
- Scope
- Type Definitions

### **Skills**

- Defining variables for given types
- Using variables in place of literals
- Be able to cleanly define multiple variables
- Get a basic understanding of scope in Go

### **Variables**

_**Basics**_

Variables in Go, like other languages, store data to be used throughout your application. There are a few basic rules that must be followed when using a variable in Go.

1. The type must be declared for any given name.
2. Once a type is declared, that name cannot be converted to another type.
3. Any function or below scoped declared variable must be used or you cannot compile.
4. Names can contain:
   - letters(any case)
   - numbers(other than the first character)
   - underscore `_` character

**Syntax of a variable:** `[var | constant] <name> <type>`

### **Declaration**

Basic variable declarations consists of naming the variable and giving it a type.

```Golang
var thing string
```

We now have a variable named `thing` that we can use to assign the value of a string.

```Golang
thing = "Oh happy days"
```

All together:

```Golang
var thing string
thing = "Oh happy days"
```

or

```Golang
var thing string = "Oh happy days"
```

or we can drop the type and **allow the assigned value to determine the type of the variable.**

```Golang
var thing = "Oh happy days"
```

---

**NOTE:**

- If the type is left off, then a value has to be assigned.
- The type will then always be the type of the value of the supplied to the right of the `=`

---

Now let's apply this to the other basic types and use them:

```Golang
package main

import "fmt"

func main() {
    var thing string
    var number int
    var floatnumber float64
    var didYouHaveToLookUpPi bool

    thing = "Look, look, look at me! I'm, I'm, I'm A tree!"
    number = 42
    floatnumber = 3.14
    didYouHaveToLookUpPi = true

    fmt.Println(thing)
    fmt.Println(number)
    fmt.Println(floatnumber)
    fmt.Println(didYouHaveToLookUpPi)
}
```

### **Short Declaration**

For the most part, in any programming language, less typing resulting in the same effect is a good thing. In Go, you have the ability to shorten the declaration of variables in a function level using the assignment operator `:=`.

**Syntax of a short variable declaration:** `<name> := <value>`

**Short variable rules:**

- A value has to be assigned
- The type will then always be the type of the value of the supplied to the right of the `:=`

Reusing the example above, we can assign and declare at the same time:

```Golang
thing := "Oh happy days"
```

The short declaration only works if there is a new name to the left of the `:=`. If you attempt to use to set a value to a variable name using `:=` again as is done in the following example:

```Golang
thing := "Oh happy days"
thing := "Not a new thing"
```

it will result in a compile error: `no new variables on left side of :=`

### **Multiple Variable Values and Declarations**

Go also supports declaring and assigning values to a variable on one line. Names get assigned types and values from left to righ on either side of the `=`.

```Golang
package main

import "fmt"

var i, j int = 1, 2 // (1)
var k, l = "abc", "def" // (2)
var (
    x,y, z = "x", 1, 2.34
    whyNotAddAnotherHere = "Firefly" // (3)
)

func main () {
    var c, python, java = true, -1, "no!" // (4)
    d, e, f, g := "c", 1, false, 1.2 // (5)

    fmt.Println(i, j, c, python, java)
    fmt.Println(d, e, f, g)
}
```

1. 2 Package level variables assigned as an `int`
2. 2 Package level variables, with the types inferred by the values
3. Mix of package level declaration types inside of `( )`
4. 3 Function level variables, with 3 different types on the right.
5. Demonstrate using the shorthand `:=` operator to assign values of various types.

### **Constants**

_**What is a constant?**_

- Constants are variables that cannot have their value changed after declaration
- Must be assigned a value when declared.

`const z = 2` Creates a constant name `z`. `z` will always have a value of 2.

---

### **Lab 1 - Declare and use Variables**

**Part 1:**
Using your own file, or a [goplay.space](https://goplay.space/)

1. Create 2 integers and print them
2. Create 2 float64 and print them
3. Create 2 strings and print them
4. Create a variable that is the sum of the two integers and print it.
5. Create a variable that is the sum of the two floats and print it.
6. Concatentate the 2 strings into a new variable and print it.

**Part 2:**

Try different variations of declaring the variables. Some ideas to try:

1. All numbers on one line
2. Everything on one line
3. Outside the main

---

### **Scope**

Scope defines where a variable is able to be accessed. In Go, a variable's scope is defined by the block where it is initially declared.

_**Blocks**_

In general, blocks are defined as code between curly braces `{ }`. There are some implicit blocks, such as at the package level block (aka package block) that are not as easy to visualize. We will cover the most common cases of blocks in this course.

```Golang
func main() {
    x := 1 // (1)
    { // (2)
        y := 2 // (3)
        fmt.Println(x)
    }
    fmt.Println(y) // (4)
}

```

1. Variable declared in the function block.
2. Anonymous block that declares the variable `y`.
3. Function block that attempts to print `y`.

_**Package Block Scope**_

The package block is anything after the import statement. We will discuss packages later on, but it is good to know that a package block can span all `.go` files that package.

```Golang
package main

import "fmt"

// Package block
var packageVariable ="Hi"
```

The `packageVariable` is declared in the package block, therefore it is able to be accessed everywhere in the package.

_**Function Block Scope**_

```Golang
package main

import "fmt"

// Package block
var packageVariable = "Hi"

func main() { // Start of the function block
    // Function block
    x := "c"
    fmt.Println(packageVariable, "-", x)
}
```

The `x` variable is only accessible to "things" within the `{}` of the main function. However, the main function can access the package block variable `packageVariable` just fine.

---

**NOTE:**

Other scope access includes loops, if statements, etc.

---

### **Type Definitions**

In Go, custom types can be defined using the `type` keyword. The type will maintain the underlying type as its type and operations, but bind to the identifier.

```Golang
// Define a solo new type.
type NewTypeName SourceType

// Define multiple new twpes together.
type (
    NewTypeName1 SourceType1
    NewTypeName2 SourceType2
)
```

```Golang
type TimeZone int
var x TImezone
x = 5
```

- A new defined type and its respective source type in type definitions are two distinct types.
- Two different defined types are always two distinct types.
- The new defined type and the source type will share the same underlying type (see below for what are underlying types), and their values can be converted to each other.
- Types can be defined within function bodies.

_**Type Alias**_

There is another type of `type` definition in Go and it is known as a `type alias` that was introduced in `go.1.9`.

The two types we covered in the `types` session, `rune` and `byte` are the only two built-in types. You can declare an `alias type` by using the `=` after the type identifier.

```Golang
type TimeZone int // (1)
type MyTimeZone = TimeZone // (2)
var y MyTimeZone
var x TimeZone

x = 5
y = 123
```

1. Type TimeZone
2. Type alias MyTimeZone

There is another concept of defined and undefined type we won't go over, but if you would like to dig deeper into type definitions and what all this means, [Go 101](https://go101.org/article/type-system-overview.html) covers a thorough understanding of types in Go.

---

## 5. **Control Structures**

### **If Statements**

If statements, also known as `if-then` or `if-then-else` statements, are how we tell a program what it should do based on some true/false condition.

In pseudo code for basic if structures in just about any language:

Pseudo code of if logic:

```Golang
if a equals b then
  do something
else if a equals c then
  do something else
else
  do some default thing
```

In Go, a similar statement would look like:

If, else if, else

```Golang
a, b, c := 1, 2, 1

if a == b {
    fmt.Println("A and B are the same!")
} else if a == c {
    fmt.Println("A and C are the same!")
} else {
    fmt.Println("Nothing is equal to A!")
}
```

You are not required to have all three conditions. One of the most common uses of `if` tends to be just the `if`, without an else, or an `else if`:

Single if statement

```Golang
a, b := 1,1

if a == b {
    fmt.Println("A and B are the same!")
}
```

\***\*Boolean Operators with `if`\*\***

Boolean operators( `&&`, `||`, `!`, etc) can be used to check multiple conditions in an if statement.

If with boolean operators

```Golang
a, b, c := 1, 1, 1
if a === b || a == c && c != b {
    fmt.Println("A has an equal!")
}
```

\***\*A Case for Parentheses\*\***

A lot of programming languages out there force you to put parenthesis around conditionals in if-statements. Golang does not allow parenthesis to occur around the **whole** conditional, but there are times you might want to put parenthesis around part of the conditional to have one operation happen earlier in the order of operations.

The following shows a case for when the placement of parenthesis can affect output.

```Golang
package main

import (
  "fmt"
)

func main() {
  a, b, c := 1, 1, 1

  if c != b && a ==b || a == c {
    fmt.Println("Tru!")
  }

  if c != b && (a == b || a == c) {
    fmt.Println("Trew!")
  }
}
```

\***\*"Short" If Statements\*\***

Short, also know as single line, if statements contains at least one variable declaration, followed by an evaluation:

Short if statements

```Golang
y := 4

if x := y % 2; x == 0 {
  fmt.Println("Even!")
}
```

In a single `if` statement:

- declares a new variable `x`
- assigns the result of `y % 2` (remainder of `4/2` is `0`, so `x` is assigned type `int` with a value of `0`)
- The x is then used after the semicolon `;` as a conditional to determine if it is `== 0`

_**Scope of `x`**_

```Golang
func main () {

  y := 4
  if x := y % 2; x == 0 {
    fmt.Println("Even")
  }
  fmt.Println(x)
}
```

In this case, the scope of `x` is not accessible to the print statement outside the `if` control structure.

\***\*If syntax options\*\***

It is possivle to write if statements all on one line. Some prefer this, others don't care. Most editors with built-in auto formatting will force multi-line statements.

The two examples from previous section could be written as:

```Golang
a, b, c := 1, 2, 1

if a == b {fmt.Println("A and B are the same!")} else if b == c{fmt.Println("no")} else { fmt.Println("No matches")}

if a == b {fmt.Println("A and B are the same!")}
```

The only real rule with syntax is that a left curly brace `{` must come before a new line character after the condition.

**Lab 1: If Statements**
**Setup**

1. Follow the instructions in the readme for the go foundations lab
2. Open and complete the `ifthen.go` file.

### **Switch Statements**

Switch statements are often used when a result of a conditional could result in many different decisions.

Although this is ok in an `if` statement:

Verbose if else-if statement

```Golang
if a == 1 {
  fmt.Println("Choice 1")
} else if a == 2 {
  fmt.Println("Choice 2")
} else if a == 3 {
  fmt.Println("Choice 3")
} else if a == 4 {
  fmt.Println("Choice 4")
}
```

The same logic could be coded more cleanly as a `switch` statement:

```Golang
switch a {
  case 1:
    fmt.Println("Choice1")
  case 2:
    fmt.Println("Choice2")
  case 3:
    fmt.Println("Choice 3")
  case 4:
    fmt.Println("Choice 4")
}
```

#### Default

Similar to the `else` in an `if-then-else` statement, there is an option to provide default behavior if none of the conditions in a `switch` block are true. With the `switch` statement, this is done with the `default` keyword:

Switch with a default option

```Golang
switch a {
  case 1:
    fmt.Println("Choice1")
  case 2:
    fmt.Println("Choice2")
  case 3:
    fmt.Println("Choice 3")
  case 4:
    fmt.Println("Choice 4")
  default:
    fmt.Println("No valid choice")
}
```

#### Short and More Complex Evaluations

As with the short `if` statements, you can have "in-line declarations" when starting with the `switch statement`.

Short switch statement

```Golang
package main

import (
  "fmt"
  "time"
)
func main() {
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  }
  default:
    fmt.Println("It's a weekday")
}
```

#### Switch with no Conditon

A `switch` statement does not require a condition. It can reference another in-scope variable or function to determine if any of the cases are true.

Switch statement with no conditions

```Golang
package main

import (
  "fmt"
  "time"
)

func main () {
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }
}
```

### **Lab 2: Switch Statement**

1. Follow the setup instructions in lab 1
2. Open `switch.go` in your editor and follow the instructions in the readme.

---

### **For - One Loop to Rule Them All!**

First, what is a loop? In programming a loop is simply a way to do repetitive work.

#### Code that should probably be in a loop

```Golang
// Print numbers 0 - 4
fmt.Println(0)
fmt.Println(1)
fmt.Println(2)
fmt.Println(3)
fmt.Println(4)
```

Instead we could use a loop to automatically repeat a block of code for us based on some condition.

Golang only has one loop and its name is `for`. We'll look at how to use the `for` loop traditionally and then explore how to apply `for` to alternative looping options you may haveseen in other languages.

#### Basic iterative for loop\*

```Golang
for i := 0; i < 5; i++ {
  fmt.Println(i)
}
```

From here, the next most common loop is traditionally a `while` loop. `while true`, `do something`

In Go, you simply replace `while` with `for` and it works the same.

#### While style loop

```Golang
run := true // (1)
for run {
  fmt.Println("Hello")
  run := false // We mark it false and the loop will stop (3)
}
```

1. Declare a `bool` called run by setting it to `true`
2. Create a `for` loop that will run until `run` is false
3. Set `run` to `false` so the loop will bot run again.

#### Continue, Break, Return and Labels

You do not need a condition in a for loop, but you need to be careful.

My first Golang infinite loop.

```Golang
for {
  // (1)
}
```

1. `true` is inferred here, and because we never toggle a false, this will run for ever.

So if we don't have access to a variable, how can we exit this loop?

##### Break & Return

Using `break`

```Golang
for {
  return
}
```

#### Continue

If you are in a loop and want to skip to the next iteration in at a particular bloc, you can use the `continue` keyword to begin the next iteration.

```Golang
func main() {
  for i := 0; i < 10; i++ {
    if x := i % 2; x == 0 {
      continue
    }
    fmt.Println(i)
  }
}
```

#### Labels

Labels can be added as part of your code before a statement. When used, they can be provided as arguments to a few key words as an indicator of where in the code to continue executing from. Labels can be passed stated after one of the following keywords to jump to a label in the code:

- break
- continue
- goto

Of the above, `goto` is the only keyword where the label is required as an argument.

#### Using Labels with `continue` and `break`

Multiple Labels on Nested loops

```Golang
func main() {
  MyLoop:
    for l := 0; l < 5; l++ {
      Derp:
        for i := 0; i < 10; i++ {
          fmt.Println("even")
          continue Derp // (1)
        }
        fmt.Println("odd")
        break MyLoop // (2)
    }
}
```

1. Skip the iteration and continue from `Derp` label
2. Break out MyLoop

#### Goto

As stated berfore, `goto` is the only keyword that a label is not optional. There are a couple of rules to keep in mind with `goto`:

- A `goto` statement transfers control to the statement with the corresponding label within the same function.
- Executing the `goto` statement must not cause any variables to come into scope that there were not already in scope at the point of the goto.
- A `goto` statement outside a block cannot jump to a label inside that block

The following examples _will cause an error_

Creates a new variable between the goto and label

```Golang
goto L // BAD
  v := 3 // New variable

L:
  someCodeCalled()
```

Cannot "goto" a child or sibling block

```Golang
if n%2 == 1 {
  goto L1
}

for n > 0 {
  f()
  n--

L1:
  f()
  n--
}
```

However, you can use `goto` to accomplish things, like creating a loop.

Reproducing a Loop with Goto

```Golang
func main() {
  i := 0
  GoHere:
    fmt.Println("Where I will go", i+1)

    if i == 10 {
      return
    }
    goto GoHere
}
```

### **Lab 3: Loops**

1. Follow the setup instruction in lab1
2. Open and complete `loop.go` in your editor.

---

### Putting it All Together: If, Switch and Loops

```Golang
package main

import "fmt"

func main() {
  for i := 0; i <= 500; i++ {
    // Is the number even?
    if x := i % 2; x == 0 {
      fmt.Printf("%d is Even \n", i)
    }

    // Do something based on what i is
    switch i {
      case 100:
        fmt.Println("400 to go!")
      case 125:
        fmt.Println("You know what? I'm done!")
        i += 300
      case 480
        fmt.Println("That's it.")
        break
        fmt.Println("Done")
    }
  }
}
```

---

## 6. **Array, Slices and Maps**

### **Theory**

- Creating and Using
  - Arrays
  - Slices
  - Maps
- Using range to loop over arrays, slices and maps

### **Objectives**

- Create an array of a given size and read/write values
- Determine the length of an array, slice, and map
- Append to a slice
- Understanding looping over strings

### **Arrays**

Arrays in Go can be defined as a fixed size collection of elements of the same type.

You can have an array of any type in Go.

Syntax: `[size]type` or `[size as int]type{comma separated literals}`

- Example 1: `[6]strings` string array with a size of 6
- Example 2: `[5]int{1, 2, 3, 4, 5}` int array with a size of 5, and default values
- Example 3: `[...]int{1,2,3,4,5}` The same as example 2, except the size will be determined based on the number of initial elements. In this case, 5.

**Quick Notes on Working with Array Variables:**

- Arrays can be declared with the same rules as any variable.
- An array's element type cannot change
- An array's size cannot change

An example of declaring an array

```Golang
  dogs := [4]string{}
```

The above will create an empty array, that can hold up to 4 strings and assign it to the variable `dogs`.

&nbsp;

#### **Indexes**

Before we can assign any values to this, we need to understand how to reference an `index` of an array. An `index` simply means the position a value is in. The first element in an array is always at `index 0` and increments by one for each item in the array.

In the `dogs` variable above, we created an array wih the size of `4`. So it has integers `0`, `1`, and `3` as indices.

To visualize this, imagine the following table is a representation of our `dogs` array:

|  Index:  | `0` | `1` | `2` | `3` |
| :------: | :-- | :-- | :-- | :-- |
| Elements | ""  | ""  | ""  | ""  |

&nbsp;

#### **Assigning Values**

##### Specific Values

To populate the elements, you need to supply the index of the element you want to store the value in.

Assigning Values by Index

```Golang
dogs := [4]string{} // Create an array with size 4
dogs[0] = "German Shepherd" // First position, index 0
dogs[3] = "Boxer" // Last position, index 3
```

This would change our table to:

|  Index:  | `0`               | `1` | `2` | `3`     |
| :------: | :---------------- | :-- | :-- | :------ |
| Elements | "German Shepherd" | ""  | ""  | "Boxer" |

As you can see, you are not required to follow any order when assigning a value. As long as the index is in the array, you can assign a value to that element at the provided index.

&nbsp;

##### Initiliazing With Values

Another option is to initialize the variable with each element assigned a value.

The dogs Array Initialized With Literals

|   Index   | `0`      | `1`        | `2`               | `3`         |
| :-------: | :------- | :--------- | :---------------- | ----------- |
| Elements: | “Collie” | “Labrador” | “German Shephard” | “Dalmatian” |

&nbsp;

#### **Reading Values**

Reading values is done much like writing values. To read a value from an array you just need to reference the index of the element you want to read the value of.

Simple Read and Print

```Golang
dogs := [4]string{"Collie", "Labrador", "German Shephard", "Dalmatian"}
fmt.Println(dogs[2]) // Would print German Shepherd
```

We can also use the value of an element to declare a new variable.

Using the value of an array to declare a new string

```Golang
dogs := [4]string{"Collie", "Labrador", "German Shephard", "Dalmatian"}
dog := dogs[2]
```

&nbsp;

##### Determining the Size of an Array

The size of an array is not always obvious. For example, it is possible to size an array based on the result of a calculation. So, if you want to determine what the size of an array is, you can utilize the built-in `len()` function.

```Golang
dogs := [4]string{"Collie", "Labrador","German Shepherd"}
size := len(dogs)
```

`len()` returns an `int` with the size of the array it is supplied.

In this example, `size` is assigned an `int` value of `4`.

#### Quick Quiz

Given the following array:

```Golang
numbers := [4]int{1,2}
```

- What type are the elements of this array?
- Without using `len()`, what is the size of the `numbers` array?
- What is the value of the element at `numbers[3]`?

&nbsp;

### **Slices**

Slices are very much like arrays, but they have a lot more flexibility. We will demonstrate the differences slices have from arrays and how these differences make slices more flexible.

You can initialize a `slice` just like `array` by leaving off the size:

Basic Slice Declaration

```Golang
dogs := []string{"Collie", "Labrador", "German Shepard", "Dalmatian"}
```

From here, you can read andwrite to the slice exactly how you would an array.

```Go
dogs[0] = "Boxer"
```

Changes the value of the element at index `0` to `"Boxer"`

#### **Creating a slice from an array**

To create a slice from an array you can use the `[low:high]` expression:

First lets define what low and high mean.

- **low**: the lowest index you want to start at. In this case, index 1.

- **high**: the highest index you want to **stop at, but not include.** You can think of this as the value at `index high - 1`. High can also be the length of the array.

Using the [low:high] expression

```Go
arr := [5]string{"a", "b", "c", "d", "e"}
x := arr[1:3] // x now is a slice with ["b", "c"]
```

This creates a slice with a length of 2, and the values of `b` and `c`. Let's use the following table to explore how that works.

| **Index**    |  0  |   1   |   2   |  3  | 4   |
| :----------- | :-: | :---: | :---: | :-: | --- |
| **Elments**  |  a  | **b** | **c** |  d  | e   |
| **low:high** |     |   1   |   3   |     |
| **Return**   |     | **b** | **c** |     |     |

You can also copy an entire array to a slice using `low:high`. When you omit `low` or `high` they defailt to the `0` and `len(array)`, respectively. Here are several methods:

Copy an entire array many different ways

```Go
arr := [5]string{"a", "b", "c", "d", "e"}

verbose := arr[o:len(arr)] // (1)

exact := arr[0:5] // (2)

omitlow := arr[:5] // (3)

omithigh := arr[0:] // (4)

omitboth := arr[:] // (5)
```

1. Using the starting index of 0 as `low` and the length of the array as `high`

2. Using the starting index of 0 as `low` and the literal `5` as `high`

3. Omitting `low` and the literal `5` as `high`

4. Using the starting index of 0 as `low` and omitting `high`

5. Omitting both `high` and `low`

#### **Sizing empty slices: Introducing make()**

Lets look at another way to create a slice. Although you can resize a slice (and we will explore how to do that shortly) you sometimes want to start out with an initial size.

For example, you could create a slice with the following syntax:

```Go
mySlice := []string{}
```

but this just creates a 0 length slice. With a 0 length slice, you have no indexes, therefore; you cannot read/write any values. (`mySlice[0]` would cause an error, because there is no index 0!)

So how can we create an empty slice that is ready to work with?

Creating an empty slice

```Go
mySlice := make([]string, 10)
```

This creates an empty slice, with 10 string type elements.

#### **Slice Functions**

There are two built-in function that you can use with slices to make common tasks easier.

##### **append()**

The `append()` built-in will allow you to easily grow a slice by adding elements.

Append Examples

```Go
package main

import (
  "fmt"
)

func main() {
  slice1 := []int{1, 2, 3}

  // Append 1 or more individual elements to the slice
  slice1 = append(slice1, 4,5)
  fmt.Println(slice1)

  slice2 := []int{6, 7, 8, 9}

  // Append combine 2 slices
  slice1 = append(slice1, slice2...)
  fmt.Println(slice1)

  slice3 := []int{5, 2, 3, 10, 11, 12, 0}

  // Select specific elements in a slice to append
  slice1 = append(slice1, slice3[3:6]...)

  fmt.Println(slice1)
}
```

The `...` following an array or slice is known as a `spread` operator. It extracts values of a given array. It is common to use this when merging slices, or passing them to `variadic functions`.

Use `append` only to append new values to a given slice to avoid undesirable results. As [this example](https://goplay.space/#HgYXczZL-Id) demonstrates, you can get unexpected results from appending to a new slice. This [goes into detail on why](https://medium.com/@Jarema./golang-slice-append-gotcha-e9020ff37374).

&nbsp;

##### **copy()**

Another method of growing a slice to copy the contents of one slice to a larger slice.

Copy to a larger slice

```Go
package main

import (
  "fmt"
)

func main() {
  slice1 := []int{1, 2, 3}
  slice2 := make([]int, 6)

  copy(slice2, slice1)

  fmt.Println(slice2) // [1 2 3 0 0 0]
}
```

`copy` is also a good method of assigning values to a newly created array.

It's worth noting that Go will let you shrink a slice as well. If you copy a larger slice to a smaller one, it will drop elements that do not fit.

```Go
package main

import (
  "fmt"
)

func main() {
  slice1 := []int{1, 2, 3, 4, 5, 6}
  slice2 := make([]int, 2)
  
  copy(slice2, slice1)

  fmt.Println(slice2) // [1, 2]
}
```

#### **Further Reading on Arrays and Slices**

For anyone interested in a deeper explanation of arrays and slices, the [official golang blog](https://blog.golang.org/) has [a great write up and additional examples](https://blog.golang.org/slices-intro).

### **Lab 1: Arrays and Slices**

Follow the instructions in foundations-labs folder for Arrays and Slices.

### **Maps**

Maps are similar to arrays and slices in that you have a collection of types. However, rather than accessing an element by an `index`, you access it by a `key` and the elements are unordered. Maps have fast read and write access compared to arrays and slices.

**Syntax for creating a map**: `map[<key_type>]<value_type>`

The **<key_type>** inside of the `[]` will determine the type of the key used to access value. The **<value_type>** following the `]` will determine the type of the value for that element. 
