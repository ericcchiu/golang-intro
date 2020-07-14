# Go Foundations

## Content

1. [Introduction](##1.-Introduction)
2. [Setting Up Golang](##2.-Setting-Up-Golang)

    2.1. [Installation](###Installation)

    2.1.1 [Windows Installation](####Windows-Installation)

    2.1.2 [Linux Installation](####Linux-Installation)

    2.2. [Installation](###Installation)

3. [Types](##3.-Types)

    3.1. [Numbers](###Numbers)

    3.2. [Strings](###Strings)

    3.3. [Boolean](###Boolean)

4. [Go Variables](##4.-Go-Variables)

    4.1. [Variable Declarations](###Declaration)

    4.2. [Short Variable Declarations](###Short-Declaration)

    4.3. [Constants](###Constants)

    4.4. [Multiple Variables](###Multiple-Variable-Values-and-Declarations)

    4.5. [Scope](###Scope)

    4.6. [Type Definitions](###Type-Definitions)

***

## __1. Introduction__

Go(Golang) was created by Google in November of 2009 originally to be an amazing procedural language for scalable cloud software. It was built to be simple and fast.

### Use Cases:

* System Admin scripts
* Web app that serves HTML pages
* File server
* Scripts that deploy code to environments

### Why Go?

* It's simple. The overarching reason behind every feature is readability.

* It compiles down to binaries so deploying is simple.

* It comes with only the basic structures and methods. Add packages only as needed. This allows the source code to remain light weight.

* It comes with an auto formatter, so spacing is not an argument on teams.
* Basic testing is built-in

* Concurrency is made easy for the developer. No thread management needed. The design of communication eliminates race conditions.

* It comes with all the things you need to make a server.

* Version 1 is completely backwards compativle, as promised.

* Interfaces allow for abstraction, easy testability (through Dependency Injection), and the Adapter Pattern.

* Consisten documentation found at [golang.org/doc](https://golang.org/doc/)

***

## __2. Setting Up Golang__

### Learning Objectives:

* Installation of Golang
* Steps to set up Golang
* Go Modes: Go Modules and GOPATH
* Parts of a Go file
* Package, Import and Declaration sections of a Go file
* Building and running a Go file

### __Installation__

#### Windows Installation

1. To install Golang on a Windows system, download and install the latest 64-bit Go set for Microsoft Windows OS: [https://golang.org/install/](https://golang.org/doc/install).
2. Follow the instructions on the Go installation program.
3. Run the Command Prompt on your computer by searching for "cmd". Open the command line and type: `"go version"`.
4. One example of the output of `go version` is:

```Linux
go version go1.14.4 windows/amd64
```

#### Linux Installation

#### Step One: Install Go Language

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

#### Step 2: Set up Go environment

Set up the Go language environment variable so that Go can be used outside the `$GOPATH`.

Add /user/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile(for a system-wide installation) or $HOME/.profile:

```linux
export PATH=$PATH:/user/local/go/bin
```

1. [Introduction](##Introductiones made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

***

## __3. Types__

### Concepts

* What is a built-in type
* The number types
* The string type
* Boolean types

### Intro

Types are how we identify a data type in Go. There are many different types in Go, such as `struct` and `interface`.

### __Numbers__

In Go there are two major number types: **Integers** and **Floating point** numbers.

#### Integers

Integers are non-decimal based numbers that can be represented using one of the following types:

* The `int` types can represent any positive or negative non-decimal number
* The `uint` types can only be positive non-decimal numbers

*Sized Types:*

* `int8`, `int16`, `int32`, and `int64`
* `uint8`, `uint16`, `uint32`, and `uint64`

The 8, 16, 32, and 64 by each type represents the maximum size of the type in bits. `int8` has a max size of 8 bits and `uint64` has a max size of 64 bits.

Examples:

* Given a 64 bit processors `int` will default to `int64`
* Given a 32 bit processors `int` will default to `int32`

#### Float

Floating point numbers are decimal numbers. Floating point numbers are represented by either `float32` or `float64`. Like `int`, the 32 and 64 represent the bit capacity of the floating point number.

There are also `complex64` and `complex128` types that fall under floating point that we will not cover here.

#### Using Number Literals

A literal is a value written as is, and not assigned to any variables.

It is very common to use the following operators with numbers when doing various calculations.

|Operator    | Description        |
|:----------:|:-------------------|
| +          | addition           |
| -          | subtraction        |
| *          | multiplication     |
| /          | division           |
| %          | remainder/Modulus  |

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

***

### __Strings__

String types are used to represent text. They are made up of a series of characters. More specifically, strings are made up of `bytes`, where there is typically one byte per character.

*__Creating Byte Literals__*

There are two ways to create string literals. A string literal can be created by surrounding the characters with double quites `"` or "back ticks" `.

*__Double Quotes__*

* Example: `"Get Going"`
* Will translate escape escape characters. For eample `\n` is translated to a newline.

*__Back Ticks__*

* **Example:** ``Get Going``
* Will not translate escape characters.
* Handy when creating literals that contain double quotes, such as when working with JSON strings.

*__Concatenated Strings__*

You often need to combine one more or more strings at some point. The quickest, most common way to do this is to use the `+` operator to combine the two strings.

**Example:** ``` "Go on " + "with your bad self." ```

There are other methods of concating strings that may be more efficient when working with large strings, or gorups of strings; however, using the `+` is perfectly fine in most cases. The other mehods will not be covered in this course, but if you are interested in some of the other methods [this is a good blog presenting](https://www.hermanschaaf.com/efficient-string-concatenation-in-go/) some benchmarking and techniques.

Alternate concatenation methods:

* Method 1: Naive Appending with ``+=``

    This is the most obvious approach to the problem. Use the concatenate operator (+=) to append each segment to the string. The thing to know about this is that strings are immutable in Go - each time a string is assigned to a variable a new address is assigned in memory to represent the new value. This is different from languages like C++ and BASIC, where a string variable can be modified in place. In Go, each time you append to the end of a string, it must create a new string and copy the contents of both the existing string and the appended string into it. As the strings you are manipulating become large this process becomes increasingly slow. In fact, this method of naive appending is O(N^2).

* Method 2: strings.Join

    The Go strings standard library package provides a Join function that takes a slice of strings and concatenates them. In our case, this requires first building a slice of strings and then calling the function on the slice. Building a slice takes time if you don’t already have one, and we therefore count this as part of the time taken to perform the concatenation in our benchmarks. The final concatenation will be O(N), but building the slice is O(N) (average) and uses O(N) memory.

    Here is an example of using ``strings.Join``:

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

* Method 3: bytes.Buffer

    The last method we’ll be evaluating is using the Buffer type from the bytes package. bytes.Buffer implements io.Writer, and can concatenate strings in O(N) time without requiring us to build a new slice.

    Here is an example of using ``bytes.Buffer``:

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

* __Working with String Literals__

    ```Golang
    package main
    imprt "fmt"
    func main() {
        fmt.Println("Part one " + "joined with " + "Part 2")
        fmt.Println("I'm a string Literal!\n", `So am I!`)
        fmt.Println("Printing a JSON String: %s \n", `{"color": "red"}`)
    }

    ```

***

### __Boolean__

Boolean values represented by either `true` or `false`.

The `bool` type can define a variable that represents either of these states.

#### Operators

There are five operators used with boolean values in Go.

|Operator|Description     |
|:------:|:---------------|
|`&&`    |And             |
|`||`    |Or              |
|`!`     |Not             |
|`==`    |Equals          |
|`==`    |Equals          |
|`!=`    |Not Equals      |

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

***

## 4. __Go Variables__

### Learning Objective:

### Variable Concepts

* Variable Declarations
* Short Variable Declarations
* Constants
* Multiple Variables
* Scope
* Type Definitions

### Skills

* Defining variables for given types
* Using variables in place of literals
* Be able to cleanly define multiple variables
* Get a basic understanding of scope in Go

## __Variables__

*__Basics__*

Variables in Go, like other languages, store data to be used throughout your application. There are a few basic rules that must be followed when using a variable in Go.

1. The type must be declared for any given name.
2. Once a type is declared, that name cannot be converted to another type.
3. Any function or  below scoped declared variable must be used or you cannot compile.
4. Names can contain:
    * letters(any case)
    * numbers(other than the first character)
    * underscore `_` character

**Syntax of a variable:** `[var | constant] <name> <type>`

### __Declaration__

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

***

**NOTE:**

* If the type is left off, then a value has to be assigned.
* The type will then always be the type of the value of the supplied to the right of the `=`

***

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

### __Short Declaration__

For the most part, in any programming language, less typing resulting in the same effect is a good thing. In Go, you have the ability to shorten the declaration of variables in a function level using the assignment operator `:=`.

**Syntax of a short variable declaration:** `<name> := <value>`

**Short variable rules:**

* A value has to be assigned
* The type will then always be the type of the value of the supplied to the right of the `:=`

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

### __Multiple Variable Values and Declarations__

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

### __Constants__

*__What is a constant?__*

* Constants are variables that cannot have their value changed after declaration
* Must be assigned a value when declared.

`const z = 2` Creates a constant name `z`. `z` will always have a value of 2.
***

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

***

### __Scope__

Scope defines where a variable is able to be accessed. In Go, a variable's scope is defined by the block where it is initially declared.

*__Blocks__*

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

*__Package Block Scope__*

The package block is anything after the import statement. We will discuss packages later on, but it is good to know that a package block can span all `.go` files that package.

```Golang
package main

import "fmt"

// Package block
var packageVariable ="Hi"
```

The `packageVariable` is declared in the package block, therefore it is able to be accessed everywhere in the package.

*__Function Block Scope__*

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
